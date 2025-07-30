//! # ANSI Escape sequences

/// Switches to the main buffer.
pub(crate) const USE_MAIN_SCREEN: &str = "\x1b[?1049l";

/// Switches to a new alternate screen buffer.
pub(crate) const USE_ALTERNATE_SCREEN: &str = "\x1b[?1049h";

/// Reset the formatting
pub(crate) const RESET_FMT: &str = "\x1b[m";

/// Invert foreground and background color
pub(crate) const REVERSE_VIDEO: &str = "\x1b[7m";

/// Move the cursor to 1:1
pub(crate) const MOVE_CURSOR_TO_START: &str = "\x1b[H";

/// DECTCTEM: Make the cursor invisible
pub(crate) const HIDE_CURSOR: &str = "\x1b[?25l";
/// DECTCTEM: Make the cursor visible
pub(crate) const SHOW_CURSOR: &str = "\x1b[?25h";

/// Clear line right of the current position of the cursor
pub(crate) const CLEAR_LINE_RIGHT_OF_CURSOR: &str = "\x1b[K";

/// Report the cursor position to the application.
pub(crate) const DEVICE_STATUS_REPORT: &str = "\x1b[6n";

/// Reposition the cursor to the end of the window
pub(crate) const REPOSITION_CURSOR_END: &str = "\x1b[999C\x1b[999B";

/*
## ANSI转义序列解析

这些十六进制序列都是ANSI转义序列，用于控制终端显示：

1. **`\x1b[?1049l`** - 切换到主缓冲区
   - `\x1b` = ESC (转义字符，ASCII 27)
   - `[?1049l` = 重置1049模式（退出交替屏幕）

2. **`\x1b[?1049h`** - 切换到新的交替屏幕缓冲区
   - `\x1b` = ESC
   - `[?1049h` = 设置1049模式（进入交替屏幕）

3. **`\x1b[m`** - 重置格式
   - `\x1b` = ESC
   - `[m` = 重置所有文本属性

4. **`\x1b[7m`** - 反显视频（前景色和背景色互换）
   - `\x1b` = ESC
   - `[7m` = 启用反显模式

5. **`\x1b[H`** - 将光标移动到1:1位置（左上角）
   - `\x1b` = ESC
   - `[H` = 光标移动到home位置(0,0)

6. **`\x1b[?25l`** - 隐藏光标
   - `\x1b` = ESC
   - `[?25l` = 重置25模式（隐藏光标）

7. **`\x1b[?25h`** - 显示光标
   - `\x1b` = ESC
   - `[?25h` = 设置25模式（显示光标）

8. **`\x1b[K`** - 清除光标右侧的当前行内容
   - `\x1b` = ESC
   - `[K` = 清除到行尾

9. **`\x1b[6n`** - 向应用程序报告光标位置
   - `\x1b` = ESC
   - `[6n` = 设备状态报告（查询光标位置）

10. **`\x1b[999C\x1b[999B`** - 将光标重新定位到窗口末尾
    - `\x1b[999C` = 向右移动999个字符
    - `\x1b[999B` = 向下移动999行
    - 实际效果是将光标移动到屏幕右下角
*/