use crate::{Solution, SolutionPair};
use std::fs::read_to_string;

///////////////////////////////////////////////////////////////////////////////

pub fn solve() -> SolutionPair {
    let file = read_to_string("inputs/day02_test.txt").expect("Failed to open input file");
    let lines: Vec<&str> = file.trim().split("\n").collect();

    let sol1: u64 = count_report(&lines) as u64;
    let sol2: u64 = count_report_with_dampener(&lines) as u64;

    (Solution::from(sol1), Solution::from(sol2))
}

fn count_report_with_dampener(lines: &Vec<&str>) -> i32 {
    let mut total_safe = 0;
    for line in lines {
        let report: Vec<i32> = line
            .split_whitespace()
            .map(|e| e.parse::<i32>().unwrap())
            .collect();

        let mut is_increasing = true;
        let mut safe = true;
        for (i, level) in report.iter().enumerate() {
            if i == 0 {
                let next = report[i + 1];
                let next_diff = next - level;
                is_increasing = next_diff > 0;
            } else {
                let prev = report[i - 1];
                let prev_diff = level - prev;

                let is_divergent = is_increasing != (prev_diff > 0);
                let is_out_of_range = prev_diff.abs() < 1 || prev_diff.abs() > 3;

                if is_divergent || is_out_of_range {
                    safe = false;
                    break;
                }
            }
        }

        if safe {
            total_safe += 1;
        }
    }

    total_safe
}

fn count_report(lines: &Vec<&str>) -> i32 {
    let mut total_safe = 0;
    for line in lines {
        let report: Vec<i32> = line
            .split_whitespace()
            .map(|e| e.parse::<i32>().unwrap())
            .collect();

        let mut is_increasing = true;
        let mut safe = true;
        for (i, level) in report.iter().enumerate() {
            if i == 0 {
                let next = report[i + 1];
                let next_diff = next - level;
                is_increasing = next_diff > 0;
            } else {
                let prev = report[i - 1];
                let prev_diff = level - prev;

                let is_divergent = is_increasing != (prev_diff > 0);
                let is_out_of_range = prev_diff.abs() < 1 || prev_diff.abs() > 3;

                if is_divergent || is_out_of_range {
                    safe = false;
                    break;
                }
            }
        }

        if safe {
            total_safe += 1;
        }
    }
    total_safe
}
