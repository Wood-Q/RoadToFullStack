use diesel::prelude::*;

#[derive(Queryable, Selectable)]
#[diesel(table_name = crate::schema::blog)]
#[diesel(check_for_backend(diesel::sqlite::Sqlite))]
pub struct Blog {
    pub id: i32,
    pub title: String,
    pub body: String,
    pub published: bool,
}

#[derive(Insertable)]
#[diesel(table_name = crate::schema::blog)]
pub struct NewBlog<'a> {
    pub title: &'a str,
    pub body: &'a str,
}
