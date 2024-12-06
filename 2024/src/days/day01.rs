use std::fs::read_to_string;

use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

pub fn solve() -> SolutionPair {
    let file = read_to_string("input/day01_test.txt").expect("Failed to open input file");
    let lines: Vec<&str> = file.trim().split("\n").collect();
    let len = lines.len();
    let mut left_vec: Vec<i32> = Vec::with_capacity(len);
    let mut right_vec: Vec<i32> = Vec::with_capacity(len);

    for line in lines {
        let v: Vec<i32> = line
            .split_whitespace()
            .map(|e| e.parse::<i32>().unwrap_or_default())
            .collect();
        left_vec.push(v[0]);
        right_vec.push(v[1]);
    }

    let sol1: u64 = 0;
    let sol2: u64 = 0;

    (Solution::from(sol1), Solution::from(sol2))
}
