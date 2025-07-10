// 声明模块
mod help;
mod models;
mod schema;

// 导入所有需要的东西
use self::help::establish_connection;
use self::models::{Blog, NewBlog}; // 导入你的模型
use crate::schema::blog::dsl::*;
use diesel::prelude::*;
use diesel::sqlite::SqliteConnection; // 导入连接函数

fn create_blog(
    conn: &mut SqliteConnection,
    new_title: &str,
    new_body: &str,
) -> Result<usize, diesel::result::Error> {
    let new_blog = NewBlog {
        title: new_title,
        body: new_body,
    };
    diesel::insert_into(blog)
        .values(&new_blog)
        .execute(conn)
}

fn find_blog(conn: &mut SqliteConnection, blog_id: i32) -> Result<Blog, diesel::result::Error> {
    blog.find(blog_id).first(conn)
}

fn get_blogs(conn: &mut SqliteConnection) -> Result<Vec<Blog>, diesel::result::Error> {
    blog.filter(published.eq(true))
        .order(id.desc())
        .load::<Blog>(conn)
}

fn update_blog(
    conn: &mut SqliteConnection,
    blog_id: i32,
    new_title: &str,
) -> Result<usize, diesel::result::Error> {
    diesel::update(blog.find(blog_id))
        .set(title.eq(new_title))
        .execute(conn)
}

fn delete_blog(conn: &mut SqliteConnection, blog_id: i32) -> Result<usize, diesel::result::Error> {
    diesel::delete(blog.find(blog_id)).execute(conn)
}

fn main() {
    let mut conn = establish_connection();
    let result = create_blog(&mut conn, "测试标题", "测试内容");
    println!("博客创建成功！");
    let find_blog = find_blog(&mut conn, 1).expect("Error finding blog");
    println!("博客标题: {}", find_blog.title);
    let find_blogs = get_blogs(&mut conn).expect("Error getting blogs");
    println!("博客列表");
    let result = update_blog(&mut conn, 1, "新标题");
    println!("博客更新成功！");
}
