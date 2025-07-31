#![allow(clippy::wildcard_imports)]

use std::fmt::{Display, Error, Write as _, format};
use std::io::{self, BufRead, BufReader, ErrorKind, Read, Seek, Write};
use std::iter::{self, repeat, successors};
use std::{fs::File, path::Path, process::Command, thread, time::Instant};

use crate::row::{HlState, Row};
use crate::{Config, Error, ansi_escape::*, syntax::Conf as SyntaxConf, sys, terminal};

//设置状态信息
macro_rules! set_status { ($editor:expr, $($arg:expr),*) => ($editor.status_msg = Some(StatusMessage::new(format!($($arg),*)))) }

const fn ctrl_key(key: u8) -> u8 {
    key & 0x1f
}
const EXIT: u8 = ctrl_key(b'Q');
const DELETE_BIS: u8 = ctrl_key(b'H');
const REFRESH_SCREEN: u8 = ctrl_key(b'L');
const SAVE: u8 = ctrl_key(b'S');
const FIND: u8 = ctrl_key(b'F');
const GOTO: u8 = ctrl_key(b'G');
const CUT: u8 = ctrl_key(b'X');
const COPY: u8 = ctrl_key(b'C');
const PASTE: u8 = ctrl_key(b'V');
const DUPLICATE: u8 = ctrl_key(b'D');
const EXECUTE: u8 = ctrl_key(b'E');
const REMOVE_LINE: u8 = ctrl_key(b'R');
const BACKSPACE: u8 = 127;

const HELP_MESSAGE: &str = "^S save | ^Q quit | ^F find | ^G go to | ^D duplicate | ^E execute | \
                            ^C copy | ^X cut | ^V paste";

//enum罗列编辑器的输入指令
enum Key {
    Arrow(AKey),
    CtrlArrow(AKey),
    PageUp,
    PageDown,
    Home,
    End,
    Delete,
    Escape,
    Char(u8),
}

enum AKey {
    Left,
    Right,
    Up,
    Down,
}

//标记光标的位置
#[derive(Default, Clone)]
struct CursorState {
    x: usize,
    y: usize,
    roff: usize,
    coff: usize,
}

impl CursorState {
    fn move_to_next_line(&mut self) {
        (self.x, self.y) = (0, self.y + 1);
    }

    fn scroll(&mut self, rx: usize, screen_rows: usize, screen_cols: usize) {
        //限制光标滚动范围，保证不超出屏幕
        //clamp函数用于将值限制在指定范围，参数是min和max
        self.roff = self
            .roff
            .clamp(self.y.saturating_sub(screen_rows.saturating_sub(1)), self.y);

        self.coff = self
            .coff
            .clamp(self.x.saturating_sub(screen_cols.saturating_sub(1)), rx);
    }
}

#[derive(Default)]
pub struct Editor {
    //提示当前模式，如果不为none，就是save，find，goto或者execute，如果是none，那么就是在编辑模式
    prompt_mode: Option<PromptMode>,
    //光标
    cursor: CursorState,
    //
    ln_pad: usize,
    //窗口宽度
    window_width: usize,
    //屏幕显示的行数
    screen_rows: usize,
    //屏幕显示的列数
    screen_cols: usize,
    //内容实际行数
    rows: Vec<Row>,
    //文件是否可修改
    dirty: bool,
    //配置
    // config:Config,
    //当退出不保存的时候提示警告的行数
    quit_times: usize,
    //文件名
    filename: Option<String>,
    //展示的状态信息
    status_msg: Option<StatusMessage>,
    //语法
    // syntax:SyntaxConf,
    n_bytes: u64,
    /// The original terminal mode. It will be restored when the `Editor`
    /// instance is dropped.
    // orig_term_mode: Option<sys::TermMode>,
    /// The copied buffer of a row
    copied_row: Vec<u8>,
}

//状态管理
struct StatusMessage {
    msg: String,
    time: Instant,
}

impl StatusMessage {
    //构造函数，返回值是自身
    fn new(msg: String) -> Self {
        Self {
            msg,
            time: Instant::now(),
        }
    }
}

fn format_size(n: u64) -> String {
    if n < 1024 {
        return format!("{n}B");
    }
    // i is the largest value such that 1024 ^ i < n
    // To find i we compute the smallest b such that n <= 1024 ^ b and subtract 1
    // from it
    let i = (64 - n.leading_zeros()).div_ceil(10) - 1;
    // Compute the size with two decimal places (rounded down) as the last two
    // digits of q This avoid float formatting reducing the binary size
    let q = 100 * n / (1024 << ((i - 1) * 10));
    format!(
        "{}.{:02}{}B",
        q / 100,
        q % 100,
        b" kMGTPEZ"[i as usize] as char
    )
}

/// `slice_find` returns the index of `needle` in slice `s` if `needle` is a
/// subslice of `s`, otherwise returns `None`.
fn slice_find<T: PartialEq>(s: &[T], needle: &[T]) -> Option<usize> {
    (0..(s.len() + 1).saturating_sub(needle.len())).find(|&i| s[i..].starts_with(needle))
}

//提示状态
enum PromptState {
    // Active contains the current buffer
    Active(String),
    // Completed contains the final string
    Completed(String),
    Cancelled,
}

/// Process a prompt keypress event and return the new state for the prompt.
fn process_prompt_keypress(mut buffer: String, key: &Key) -> PromptState {
    //允许忽略匹配不完整的枚举变量
    #[allow(clippy::wildcard_enum_match_arm)]
    match key {
        //如果按下回车键，则返回Completed状态
        Key::Char(b'\r') => return PromptState::Completed(buffer),
        //如果按下ESC键或者Q键，则返回Cancelled状态
        Key::Escape | Key::Char(EXIT) => return PromptState::Cancelled,
        //如果按下退格键或者删除键，则删除缓冲区中的最后一个字符
        Key::Char(BACKSPACE | DELETE_BIS) => _ = buffer.pop(),
        //如果按下的是可打印字符，则将字符添加到缓冲区中
        Key::Char(c @ 0..=126) if !c.is_ascii_control() => buffer.push(*c as char),
        //如果按下的是其他字符，则不进行任何操作
        // No-op
        _ => (),
    }
    PromptState::Active(buffer)
}

//提示模式
enum PromptMode {
    Save(String),
    Find(String, CursorState, Option<usize>),
    GoTo(String),
    Execute(String),
}

impl PromptMode {
    fn status_msg(&self) -> String {
        match self {
            Self::Save(buffer) => format!("Save: {}", buffer),
            Self::Find(buffer, cursor, pos) => {
                format!("Find: {} | {} | {}", buffer, cursor.y, pos.unwrap_or(0))
            }
            Self::GoTo(buffer) => format!("Go to: {}", buffer),
            Self::Execute(buffer) => format!("Command To Execute: {}", buffer),
        }
    }

    fn process_keypress(self, ed: &mut Editor, key: &Key) -> Result<Option<Self>, Error> {
        ed.status_msg = None;
        match self {
            Self::Save(b) => match process_prompt_keypress(b, key) {
                PromptState::Active(b) => return Ok(Some(Self::Save(b))),
                PromptState::Cancelled => set_status!(ed, "Save aborted"),
                PromptState::Completed(file_name)=>ed.save_as(file_name)?,
            },
        }
    }
}
