use diesel::prelude::*;
use diesel::sqlite::SqliteConnection; // 导入 SqliteConnection
use dotenv::dotenv;
use std::env;

pub fn establish_connection() -> SqliteConnection {
    // 返回类型是 SqliteConnection
    dotenv().ok();

    let database_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    SqliteConnection::establish(&database_url) // 使用 SqliteConnection::establish
        .unwrap_or_else(|_| panic!("Error connecting to {}", database_url))
}
