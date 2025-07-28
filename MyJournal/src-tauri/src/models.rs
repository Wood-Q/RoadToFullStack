use diesel::prelude::*;

#[derive(Queryable,Selectable)]
#[diesel(table_name = crate::schema::journal)] // 关联到 schema.rs 中定义的 posts 表
#[diesel(check_for_backend(diesel::sqlite::Sqlite))]
pub struct Journal {
    pub id: i32,
    pub title: String,
    pub body: String,
}

#[derive(Insertable)] // 派生宏，让 struct 可以被插入到数据库
#[diesel(table_name = crate::schema::journal)] // 关联到 posts 表
pub struct NewJournal<'a> {
    pub title: &'a str,
    pub body: &'a str,
}