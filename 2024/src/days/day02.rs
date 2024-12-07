use crate::{Solution, SolutionPair};
use std::fs::read_to_string;

///////////////////////////////////////////////////////////////////////////////

pub fn solve() -> SolutionPair {
    let file = read_to_string("inputs/day02.txt").expect("Failed to open input file");
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
        let mut unsafe_indexes = (0, 0);
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
                    unsafe_indexes = (i - 1, i);
                    break;
                }
            }
        }

        if safe {
            total_safe += 1;
        } else {
            // println!("{:?} , unsafe_indexes: {:?}", report, unsafe_indexes);
            let mut report_no_prev = report.clone();
            report_no_prev.remove(unsafe_indexes.0);
            let mut report_no_cur = report.clone();
            report_no_cur.remove(unsafe_indexes.1);
            // println!("{:?} {:?}", report_no_prev, report_no_cur);

            let report_no_prev_prev = {
                if unsafe_indexes.0 == 1 {
                    let mut t = report.clone();
                    t.remove(0);
                    t
                } else {
                    Vec::new()
                }
            };

            let left = is_really_safe(&report_no_prev);
            let right = is_really_safe(&report_no_cur);
            let special = is_really_safe(&report_no_prev_prev);
            // println!(
            //     "report_no_prev_prev?: {} report_no_prev?: {} report_no_cur?: {}\n",
            //     special, left, right
            // );

            if left || right || special {
                total_safe += 1;
            }
        }
    }

    total_safe
}

fn is_really_safe(report: &Vec<i32>) -> bool {
    if report.len() == 0 {
        return false;
    }
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
    safe
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
