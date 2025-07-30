#[derive(Debug)]
pub enum Error{
    //处理常见IO错误元组
    Io(std::io::Error),
    //处理格式化错误
    Fmt(std::fmt::Error),
    //处理光标位置错误
    CursorPosition,
    //处理过多启动参数错误
    TooManyArguments(usize),
    //处理不认识的选项的错误
    UnrecognizedOption(String),
    //处理配置错误
    Config(std::path::PathBuf,usize,String)
}

impl From<std::io::Error> for Error{
    fn from(value: std::io::Error) -> Self {
        Self::Io(value)
    }
}

impl From<std::fmt::Error> for Error{
    fn from(value: std::fmt::Error) -> Self {
        Self::Fmt(value)
    }
}