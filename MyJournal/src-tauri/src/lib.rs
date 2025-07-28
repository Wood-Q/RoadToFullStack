// Learn more about Tauri commands at https://tauri.app/develop/calling-rust/

mod help;
mod models;
mod schema;
use self::help::establish_connection;
use self::models::NewJournal; // 导入你的模型
use crate::schema::journal::dsl::*;
use diesel::prelude::*;

#[tauri::command]
// 1. 将返回类型修改为 Result<String, String>，代表 成功时返回字符串，失败时也返回字符串
fn create_journal(new_title: &str, new_body: &str) -> Result<String, String> {
    let mut conn = establish_connection();
    let new_journal = NewJournal {
        title: new_title,
        body: new_body,
    };

    // 2. 不要使用 unwrap()，而是将 Result 直接返回
    diesel::insert_into(journal)
        .values(&new_journal)
        .execute(&mut conn)
        .map(|_| "success".to_string()) // 如果成功 (Ok)，返回成功的消息
        .map_err(|e| e.to_string()) // 如果失败 (Err)，将错误转换为字符串并返回
}

#[tauri::command]
fn greet(name: &str) -> String {
    format!("Hello, {}! You've been greeted from Rust!", name)
}

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    tauri::Builder::default()
        .plugin(tauri_plugin_opener::init())
        .invoke_handler(tauri::generate_handler![greet, create_journal])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
