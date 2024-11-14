use std::cmp::PartialEq;
use std::collections::HashSet;
use std::fmt::{self, Display, Formatter};
use std::io::{stdout, Write};
use std::thread::sleep;
use std::time::Duration;

#[derive(Clone, Copy, PartialEq)]
pub enum Color {
    Red,
    Orange,
    Yellow,
    Green,
    Blue,
    Indigo,
    Violet,
    Reset,
}

impl Display for Color {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        if *self == Color::Reset {
            return write!(f, "\x1b[0m");
        }

        let code = match self {
            Color::Red => "255;0;0",
            Color::Orange => "255;215;0",
            Color::Yellow => "255;255;0",
            Color::Green => "0;255;0",
            Color::Blue => "0;0;255",
            Color::Indigo => "75;0;130",
            Color::Violet => "127;0;255",
            _ => "",
        };

        write!(f, "\x1b[48;2;{}m", code)
    }
}

pub fn render_rainbow() {
    let array: [[i32; 7]; 2] = [
        [1, 2, 3, 4, 5, 6, 7],
        [1, 2, 3, 4, 5, 6, 7],
    ];

    let mut colored_indexes = HashSet::<usize>::new();
    loop {
        for color_index in 0..14 {
            clear_console();
            if color_index >= 7 {
                colored_indexes.remove(&(color_index % 7));
            } else {
                colored_indexes.insert(color_index);
            }
            for (y, row) in array.iter().enumerate() {
                for (i, &num) in row.iter().enumerate() {
                    let cell_index = y * row.len() + i;
                    let mut color = Color::Reset;
                    if colored_indexes.contains(&cell_index) {
                        color = match cell_index {
                            0 => Color::Red,
                            1 => Color::Orange,
                            2 => Color::Yellow,
                            3 => Color::Green,
                            4 => Color::Blue,
                            5 => Color::Indigo,
                            6 => Color::Violet,
                            _ => Color::Reset,
                        }
                    }

                    print!("{}{}", color, num);
                    print!("{}", Color::Reset);
                }
                println!();
            }
            sleep(Duration::from_secs(1));
        }
    }
}

pub fn clear_console() {
    print!("\x1B[2J\x1B[1;1H");
    stdout().flush().unwrap();
}
