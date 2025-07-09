use diesel::prelude::*;
use diesel::sqlite::SqliteConnection;
use diesel_migrations::{EmbeddedMigrations, MigrationHarness, embed_migrations};

// 嵌入迁移文件
pub const MIGRATIONS: EmbeddedMigrations = embed_migrations!("migrations");

pub mod schema {
    use diesel::table;

    table! {
        users{
            id -> Integer,
            name -> Text,
            email -> Text,
            password -> Text,
        }
    }
}

use schema::users;
#[derive(Queryable, Selectable, Debug)]
#[diesel(table_name = users)]
#[diesel(check_for_backend(diesel::sqlite::Sqlite))]
pub struct User {
    pub id: i32,
    pub name: String,
    pub email: String,
    pub password: String,
}

#[derive(Insertable)]
#[diesel(table_name = users)]
pub struct NewUser<'a> {
    pub name: &'a str,
    pub email: &'a str,
    pub password: &'a str,
}

#[derive(AsChangeset)]
#[diesel(table_name = users)]
pub struct UpdateUser<'a> {
    pub name: Option<&'a str>,
    pub email: Option<&'a str>,
}

// 建立数据库连接
pub fn establish_connection() -> SqliteConnection {
    let database_url = "test.db";
    SqliteConnection::establish(database_url)
        .unwrap_or_else(|_| panic!("Error connecting to {}", database_url))
}

// 运行迁移
pub fn run_migrations(connection: &mut SqliteConnection) {
    connection
        .run_pending_migrations(MIGRATIONS)
        .expect("迁移失败");
}

fn main() {
    let mut connection = establish_connection();

    // 运行迁移（创建表）
    run_migrations(&mut connection);
    println!("数据库迁移完成");

    println!("=====开始创建用户=====");
    let new_user = create_user(&mut connection, "张三", "123456@qq.com", "password123");
    println!("创建用户成功: {:?}", new_user);
}

fn create_user(
    conn: &mut SqliteConnection,
    user_name: &str,
    user_email: &str,
    user_password: &str,
) -> User {
    use schema::users::dsl::*;

    let new_user = NewUser {
        name: user_name,
        email: user_email,
        password: user_password,
    };

    // SQLite 不支持 RETURNING，所以我们先插入，然后查询最后插入的记录
    diesel::insert_into(users)
        .values(&new_user)
        .execute(conn)
        .expect("插入用户失败");

    // 查询刚插入的用户
    users
        .order(id.desc())
        .first(conn)
        .expect("查询刚创建的用户失败")
}
