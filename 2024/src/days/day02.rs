use crate::{Solution, SolutionPair};
use std::fs::read_to_string;

///////////////////////////////////////////////////////////////////////////////

pub fn solve() -> SolutionPair {
    let file = read_to_string("input/day02.txt").expect("Failed to open input file");
    let lines: Vec<&str> = file.trim().split("\n").collect();
    let mut total_safe = 0;
    for line in lines {
        let report: Vec<i32> = line
            .split_whitespace()
            .map(|e| e.parse::<i32>().unwrap())
            .collect();

        let mut is_increasing = true;
        let mut safe = true;
        for (i, level) in report.iter().enumerate() {
            // println!("{} {}", i, level);
            if i == 0 {
                // println!("compare with next element");
                let next = report[i + 1];
                let next_diff = next - level;
                is_increasing = next_diff > 0;

                // Check difference
                if next_diff.abs() < 1 || next_diff.abs() > 3 {
                    safe = false;
                    break;
                }
            } else if i == report.len() - 1 {
                // println!("compare with previous element")
                let prev = report[i - 1];
                let prev_diff = level - prev;
                let is_still_increasing = prev_diff > 0;

                // Check increasing
                if is_increasing != is_still_increasing {
                    safe = false;
                    break;
                }

                // Check difference
                if prev_diff.abs() < 1 || prev_diff.abs() > 3 {
                    safe = false;
                    break;
                }
            } else {
                // println!("compare with next and previous element");
                let next = report[i + 1];
                let prev = report[i - 1];
                let next_diff = next - level;
                let prev_diff = level - prev;
                // Check increasing
                let is_still_increasing = prev_diff > 0;
                if is_increasing != is_still_increasing {
                    safe = false;
                    break;
                }

                // Check difference
                if next_diff.abs() < 1 || next_diff.abs() > 3 {
                    safe = false;
                    break;
                }
                if prev_diff.abs() < 1 || prev_diff.abs() > 3 {
                    safe = false;
                    break;
                }
            }
        }
        if safe {
            total_safe += 1;
        }
    }

    let sol1: u64 = total_safe as u64;
    let sol2: u64 = 0;

    (Solution::from(sol1), Solution::from(sol2))
}
