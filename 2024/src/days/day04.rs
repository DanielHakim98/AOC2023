use std::{collections::HashMap, fs::read_to_string};

use crate::{Solution, SolutionPair};

pub fn solve() -> SolutionPair {
    let file = read_to_string("inputs/day04.txt").expect("Failed to open input file");

    let sol1: u64 = part1(&file) as u64;
    let sol2: u64 = 0 as u64;

    (Solution::from(sol1), Solution::from(sol2))
}

fn part1(file: &str) -> i32 {
    let mut coordinates: HashMap<(usize, usize), char> = HashMap::new();
    for (i, line) in file.trim().split("\n").enumerate() {
        for (j, ch) in line.chars().enumerate() {
            coordinates.insert((i, j), ch);
        }
    }
    // * To be honest, I know the rough idea to solve it, but I don't know how to do it in Rust, so I end up uses ChatGPT :(
    let get_safe = |row: usize, col: usize| coordinates.get(&(row, col)).unwrap_or(&' ');
    let directions: [(&str, Box<dyn Fn(usize, usize, char) -> String>); 8] = [
        (
            "horizontal right",
            Box::new(|row, col, ch| {
                format!(
                    "{}{}{}{}",
                    ch,
                    get_safe(row, col + 1),
                    get_safe(row, col + 2),
                    get_safe(row, col + 3)
                )
            }),
        ),
        (
            "horizontal left",
            Box::new(|row, col, ch| {
                format!(
                    "{}{}{}{}",
                    ch,
                    get_safe(row, col - 1),
                    get_safe(row, col - 2),
                    get_safe(row, col - 3)
                )
            }),
        ),
        (
            "vertical down",
            Box::new(|row, col, ch| {
                format!(
                    "{}{}{}{}",
                    ch,
                    get_safe(row + 1, col),
                    get_safe(row + 2, col),
                    get_safe(row + 3, col)
                )
            }),
        ),
        (
            "vertical up",
            Box::new(|row, col, ch| {
                format!(
                    "{}{}{}{}",
                    ch,
                    get_safe(row - 1, col),
                    get_safe(row - 2, col),
                    get_safe(row - 3, col)
                )
            }),
        ),
        (
            "diagonal down right",
            Box::new(|row, col, ch| {
                format!(
                    "{}{}{}{}",
                    ch,
                    get_safe(row + 1, col + 1),
                    get_safe(row + 2, col + 2),
                    get_safe(row + 3, col + 3)
                )
            }),
        ),
        (
            "diagonal down left",
            Box::new(|row, col, ch| {
                format!(
                    "{}{}{}{}",
                    ch,
                    get_safe(row + 1, col - 1),
                    get_safe(row + 2, col - 2),
                    get_safe(row + 3, col - 3)
                )
            }),
        ),
        (
            "diagonal up left",
            Box::new(|row, col, ch| {
                format!(
                    "{}{}{}{}",
                    ch,
                    get_safe(row - 1, col - 1),
                    get_safe(row - 2, col - 2),
                    get_safe(row - 3, col - 3)
                )
            }),
        ),
        (
            "diagonal up right",
            Box::new(|row, col, ch| {
                format!(
                    "{}{}{}{}",
                    ch,
                    get_safe(row - 1, col + 1),
                    get_safe(row - 2, col + 2),
                    get_safe(row - 3, col + 3)
                )
            }),
        ),
    ];

    let mut total = 0;
    for (&(row, col), &ch) in &coordinates {
        for &(_, ref closure) in &directions {
            let result = closure(row, col, ch);
            if &result == "XMAS" {
                total += 1;
            }
        }
    }
    total
}
