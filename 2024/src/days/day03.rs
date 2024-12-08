use crate::{Solution, SolutionPair};
use regex::Regex;
use std::fs::read_to_string;

///////////////////////////////////////////////////////////////////////////////

pub fn solve() -> SolutionPair {
    let re = Regex::new(r"mul\([0-9]{1,3},[0-9]{1,3}\)").unwrap();
    let file = read_to_string("inputs/day03.txt").expect("Failed to open input file");
    let lines: Vec<&str> = file.trim().split("\n").collect();

    let sol1: u64 = mull_it_over_part_1(&re, &lines) as u64;
    let sol2: u64 = 0 as u64;

    (Solution::from(sol1), Solution::from(sol2))
}

fn mull_it_over_part_1(re: &Regex, lines: &Vec<&str>) -> i32 {
    let mut total = 0;
    for line in lines {
        for m in re.captures_iter(line) {
            total += match m.get(0) {
                Some(v) => v
                    .as_str()
                    .replace("mul", "")
                    .replace("(", "")
                    .replace(")", "")
                    .split(",")
                    .map(|e| e.parse::<i32>().unwrap())
                    .fold(1, |acc, x| acc * x),
                _ => 0,
            }
        }
    }
    total
}
