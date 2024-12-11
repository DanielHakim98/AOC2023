use crate::{Solution, SolutionPair};
use std::fs::read_to_string;

pub fn solve() -> SolutionPair {
    let file = read_to_string("inputs/day05_test.txt").expect("Failed to open input file");
    println!("{}", file);
    let sol1: u64 = 0 as u64;
    let sol2: u64 = 0 as u64;

    (Solution::from(sol1), Solution::from(sol2))
}
