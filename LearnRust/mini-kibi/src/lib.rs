//! # Mini-Kibi

pub use crate::error::Error;

pub mod ansi_escape;
// mod config;
// mod editor;
mod error;
// mod row;
// mod syntax;
mod terminal;

// #[cfg_attr(unix, path = "unix.rs")]
// #[cfg_attr(target_os = "wasi", path = "wasi.rs")]
// mod sys;

// #[cfg(any(unix, target_os = "wasi"))] mod xdg;
