use std::io::{self, BufRead, Read, Write};

use crate::{Error, ansi_escape::DEVICE_STATUS_REPORT, ansi_escape::REPOSITION_CURSOR_END};

//有些不适配获取窗口大小的，通过光标移动来实现
pub fn get_window_size_using_cursor() -> Result<(usize, usize), Error> {
    let mut stdin = io::stdin().lock();
    print!("{REPOSITION_CURSOR_END}{DEVICE_STATUS_REPORT}");
    io::stdout().flush()?;
    let mut prefix_buffer = [0u8; 2];
    stdin.read_exact(&mut prefix_buffer)?;
    if prefix_buffer != [b'\x1b', b'['] {
        return Err(Error::CursorPosition);
    }
    Ok((read_value_until(b';')?, read_value_until(b'R')?))
}

//不断读取数据，直到读到stop_byte为止
pub fn read_value_until<T: std::str::FromStr>(stop_byte: u8) -> Result<T, Error> {
    //数据缓冲区
    let mut buf = Vec::new();
    //直到读到stop_byte为止
    io::stdin().lock().read_until(stop_byte, &mut buf)?;
    //从缓冲区中删除stop_byte
    buf.pop()
        .filter(|u| *u == stop_byte)
        .ok_or(Error::CursorPosition)?;
    //将缓冲区转换为字符串，并解析为T类型
    std::str::from_utf8(&buf)
        .or(Err(Error::CursorPosition))?
        .parse()
        .or(Err(Error::CursorPosition))
}
