use crate::{Solution, SolutionPair};
use regex::Regex;
use std::fs::read_to_string;

///////////////////////////////////////////////////////////////////////////////

pub fn solve() -> SolutionPair {
    let re = Regex::new(r"mul\([0-9]{1,3},[0-9]{1,3}\)").unwrap();
    let file = read_to_string("inputs/day03.txt").expect("Failed to open input file");
    let lines: Vec<&str> = file.trim().split("\n").collect();

    let sol1: u64 = mull_it_over_part_1(&re, &lines) as u64;
    let sol2: u64 = mull_it_over_part_2(&re, &file.trim()) as u64;

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

// At the end, I actually could not solve this, I thought of developing custom parser
// to solve this problem but lack of knowledge and theory on intepreter, then I found this
// video form 'CS Jackie' => https://www.youtube.com/watch?v=8_o4625JK50
// The solution is easily understandable but I didn't profile the soluttion in term of time & space
// complexity.
fn mull_it_over_part_2(re: &Regex, content: &str) -> i32 {
    let mut total = 0;
    for item in content.split("do()") {
        for ele in item.split("don't()") {
            for m in re.captures_iter(ele) {
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
            // From what I understand, the idea is to always extract the first element after splitting
            // by "don't()", becausse that element is the part where the mul(...) is still valid before
            // "don't() appear and make the rest of mul(...) invalid".
            break;
        }
    }
    total
}
