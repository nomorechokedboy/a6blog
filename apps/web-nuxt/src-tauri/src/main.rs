#![cfg_attr(
    all(not(debug_assertions), target_os = "windows"),
    windows_subsystem = "windows"
)]

pub fn main() {
    web_nuxt::AppBuilder::new().run()
}
