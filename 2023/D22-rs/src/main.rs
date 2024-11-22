use std::cmp::{max, min};
use std::fmt::{Debug, Formatter};
use std::fs::File;
use std::io::{BufRead, BufReader, Lines};
use std::path::Path;

#[derive(Debug)]
struct Vec3<T> {
    x: T, y: T, z: T,
}

impl From<&str> for Vec3<u32> {
    fn from(value: &str) -> Self {
        let parsed: [u32; 3] = value
            .split(",")
            .map(|s| s.parse::<u32>().expect("PARSE FAILED"))
            .collect::<Vec<u32>>()
            .try_into().expect("TRY FAILED");

        Vec3 {
            x: parsed[0], y: parsed[1], z: parsed[2],
        }
    }
}

struct Brick<T> {
    min_x: T,
    max_x: T,
    min_y: T,
    max_y: T,
    min_z: T,
    max_z: T,
}

impl Brick<u32> {
    fn from_corners(a: Vec3<u32>, b: Vec3<u32>) -> Self {
        Brick {
            min_x: min(a.x, b.x),
            max_x: max(a.x, b.x),
            min_y: min(a.y, b.y),
            max_y: max(a.y, b.y),
            min_z: min(a.z, b.z),
            max_z: max(a.z, b.z),
        }
    }

    pub fn is_intersecting(&self, other: &Brick<u32>) -> bool {
        return (
            self.min_x <= other.max_x &&
            self.max_x >= other.min_x &&
            self.min_y <= other.max_y &&
            self.max_y >= other.min_y &&
            self.min_z <= other.max_z &&
            self.max_z >= other.min_z
        );
    }

    pub fn drop(&self) -> Option<Brick<u32>> {
        if self.min_z == 1 {
            return None
        }

        Some(Brick {
            min_x: self.min_x,
            max_x: self.max_x,
            min_y: self.min_y,
            max_y: self.max_y,
            min_z: self.min_z - 1,
            max_z: self.max_z - 1,
        })
    }
}

impl Debug for Brick<u32> {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        // f.debug_struct("Brick")
        //  .field("x", &self.x)
        //  .field("y", &self.y)
        //  .finish()
        write!(f, "Brick [{},{},{}~{},{},{}]", self.min_x, self.min_y, self.min_z, self.max_x, self.max_y, self.max_z)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_is_intersecting() {
        let test = true;

        let base_box = &Brick {
            min_x: 0,
            max_x: 3,
            min_y: 0,
            max_y: 3,
            min_z: 0,
            max_z: 3,
        };

        let box_intersecting = &Brick {
            min_x: 1,
            max_x: 2,
            min_y: 1,
            max_y: 2,
            min_z: 1,
            max_z: 2,
        };

        let box_non_intersecting = &Brick {
            min_x: 10,
            max_x: 11,
            min_y: 10,
            max_y: 11,
            min_z: 10,
            max_z: 11,
        };

        let box_single_axis_intersecting = &Brick {
            min_x: 3,
            max_x: 4,
            min_y: 3,
            max_y: 4,
            min_z: 2,
            max_z: 3,
        };

        assert_eq!(true, base_box.is_intersecting(box_intersecting));
        assert_eq!(false, base_box.is_intersecting(box_non_intersecting));
        assert_eq!(true, base_box.is_intersecting(box_single_axis_intersecting));
    }
}

fn shake_down(tower: &mut Vec<Brick<u32>>) {
    for (brick_index, brick) in tower.iter().enumerate() {
        while let Some(next_brick) = brick.drop() {
            let mut can_drop = true;
            for (brick_to_check_index, brick_to_check) in tower.iter().enumerate() {
                if brick_index == brick_to_check_index {
                    continue
                }
                if next_brick.is_intersecting(brick_to_check) {
                    can_drop = false;
                    break;
                }
            }

            if can_drop {
                tower[brick_index] = next_brick;
            }
        }
    }
}

fn main() {
    let file = File::open(Path::new("debug.txt")).expect("FILE ERROR");
    let lines = BufReader::new(file).lines() ;

    let mut bricks: Vec<Brick<u32>> = vec![];

    for line in lines.flatten() {
        let (start, finish) = line
            .split_once("~")
            .map(|(a, b)| (Vec3::from(a), Vec3::from(b))).expect("PARSE FAIELD");

        bricks.push(Brick::from_corners(start, finish));
    }

    for brick in bricks {
        println!("{:?}", brick);
    }
}
