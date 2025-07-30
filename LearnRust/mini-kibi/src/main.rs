//这里Result的()表示成功不需要返回任何值

use mini_kibi::Error;

fn main() -> Result<(), Error> {
    //读取启动参数
    let mut args = std::env::args();
    match (args.nth(1), args.len()) {
        //Some表示这个值存在，并且为内部填充的值
        (Some(arg), 0) if arg == "--version" => {
            println!("这个编辑器版本为{}", env!("Editor_Version"))
        }
        (Some(arg), 0) if arg.starts_with("-") => return Err(Error::UnrecognizedOption(arg)),
        //TODO 如果输入的第二个参数是文件名，那么启动编辑器
        // (file_name,0),
        (_, n_remaining_args) => return Err(Error::TooManyArguments(n_remaining_args + 1)),
    }
    Ok(())
}
