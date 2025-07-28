// @generated automatically by Diesel CLI.

diesel::table! {
    journal (id) {
        id -> Integer,
        title -> Text,
        body -> Text,
    }
}
