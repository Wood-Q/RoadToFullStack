use std::env;
use std::process;

use minigrep::Config;
use minigrep::run;

fn main() {
    let args: Vec<String> = env::args().collect();
    // let query = &args[1];
    // let file_path = &args[2];
    let config = Config::build(&args).unwrap_or_else(|err| {
        eprintln!("Problem parsing arguments: {}", err);
        process::exit(1);
    });
    // println!("The query is {}", config.query);
    // println!("The file_path is {}", config.file_path);
    if let Err(e) = run(config) {
        eprintln!("Application error: {}", e);
        process::exit(1);
    }
}
