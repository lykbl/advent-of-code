use std::thread::sleep;
use std::time::Duration;
use sysinfo::System;

mod modules;

trait Module {
    fn connect<'a>(&mut self, to_connect: &'a dyn Module) -> ();
}

struct SimpleModule<'a> {
    connected: Option<&'a dyn Module>,
}

impl<'a> Module for SimpleModule<'a> {
    fn connect<'b>(&mut self, to_connect: &'b dyn Module) -> () {
        self.connected = Some(to_connect);
    }
}

fn main() {
    println!("Hello world");
}
