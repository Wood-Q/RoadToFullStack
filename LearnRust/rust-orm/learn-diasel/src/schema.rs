// @generated automatically by Diesel CLI.

diesel::table! {
    blog (id) {
        id -> Integer,
        title -> Text,
        body -> Text,
        published -> Bool,
    }
}
