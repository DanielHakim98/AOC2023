use crate::{Solution, SolutionPair};
use regex::Regex;
use std::fs::read_to_string;

///////////////////////////////////////////////////////////////////////////////

pub fn solve() -> SolutionPair {
    let re = Regex::new(r"mul\([0-9]{1,3},[0-9]{1,3}\)").unwrap();
    let file = read_to_string("inputs/day03_test.txt").expect("Failed to open input file");
    let lines: Vec<&str> = file.trim().split("\n").collect();

    let sol1: u64 = mull_it_over_part_1(&re, &lines) as u64;
    let sol2: u64 = mull_it_over_part_2(&file.trim()) as u64;

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

fn mull_it_over_part_2(content: &str) -> i32 {
    let len = content.len();
    let mut i = 0;
    let mut enable = true;
    while i < len {
        let c = content.chars().nth(i).unwrap_or_default();

        match c {
            'd' => {
                if let Some(next) = content.get(i..i + 4) {
                    if next == "do()" {
                        if !enable {
                            enable = true
                        }
                        i += 4;
                        continue;
                    }
                }
                if let Some(next) = content.get(i..i + 7) {
                    if next == "don't()" {
                        if enable {
                            enable = false
                        }
                        i += 7;
                        continue;
                    }
                }
            }
            'm' => {
                if let Some(next) = content.get(i..i + 4) {
                    if next == "mul(" {
                        i += 4;
                        let mut j = i;
                        let mut char = content.chars().nth(j).unwrap_or_default();
                        while char != ',' {
                            j += 1;
                            char = content.chars().nth(j).unwrap_or_default();
                        }
                        if let Some(nums) = content.get(i..j) {
                            let first = nums.parse::<i32>().unwrap_or(1);
                            println!("first: {first} i: {i} j: {j}");
                        }

                        i = j + 1;
                        println!("i:{i}");
                        let mut k = i;
                        let mut char = content.chars().nth(k).unwrap_or_default();
                        while char != ')' {
                            k += 1;
                            char = content.chars().nth(k).unwrap_or_default();
                        }
                        if let Some(nums) = content.get(i..k) {
                            println!("nums:{nums}\n");
                        }
                        continue;
                    }
                }
            }
            _ => {}
        }

        i += 1;
    }

    0
}
