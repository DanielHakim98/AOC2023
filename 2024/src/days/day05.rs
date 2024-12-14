use crate::{Solution, SolutionPair};
use std::{collections::HashMap, fs::read_to_string};

pub fn solve() -> SolutionPair {
    let file = read_to_string("inputs/day05.txt").expect("Failed to open input file");

    let sol1: u64 = part1(&file) as u64;
    let sol2: u64 = part2(&file) as u64;

    (Solution::from(sol1), Solution::from(sol2))
}

// To be honest, I'm struggling with this problem,
// I think it's because previous problem, I treat it as topological problem and had no idea
// how to handle cyclic graph like
/*
    1 | 2
    3 | 4
    4 | 1

    1,2,3,4,5 <- this should be invalid by default
*/
// At the end, I just refer to a video by William Y. Feng: https://www.youtube.com/watch?v=LA4RiCDPUlI
// which shows this problem can also be treated as a sorting problem. (despite being less efficient).
fn part2(file: &str) -> i32 {
    let mut rules: Vec<(i32, i32)> = Vec::new();
    let mut total = 0;
    for line in file.trim().split("\n") {
        if line.contains("|") {
            rules.push(
                line.split("|")
                    .map(|e| e.parse::<i32>().unwrap_or_default())
                    .collect::<Vec<_>>()
                    .try_into()
                    .map(|arr: [i32; 2]| (arr[0], arr[1]))
                    .unwrap_or_default(),
            );

            continue;
        }

        if line == "" {
            continue;
        }

        let line_vec: Vec<i32> = line
            .split(",")
            .map(|e| e.parse::<i32>().unwrap_or_default())
            .collect();

        let update: HashMap<i32, usize> =
            line_vec
                .into_iter()
                .enumerate()
                .fold(HashMap::new(), |mut acc, ele| {
                    acc.insert(ele.1, ele.0);
                    acc
                });

        let mut is_right_order = true;
        for (left, right) in &rules {
            let contains_left = update.contains_key(left);
            let contains_right = update.contains_key(right);
            if contains_left && contains_right && update[&left] > update[&right] {
                is_right_order = false;
                break;
            }
        }

        if !is_right_order {
            let mut update: Vec<i32> = line
                .split(",")
                .map(|e| e.parse::<i32>().unwrap_or_default())
                .collect();
            let mut not_sorted = true;
            while not_sorted {
                let mut sorted = true;
                for i in 0..update.len() - 1 {
                    let left = update[i];
                    let right = update[i + 1];
                    if rules.contains(&(right, left)) {
                        sorted = false;
                        let temp = update[i];
                        update[i] = update[i + 1];
                        update[i + 1] = temp;
                    }
                }

                not_sorted = !sorted;
            }
            total += update[update.len() / 2];
        }
    }
    total
}

fn part1(file: &str) -> i32 {
    let mut rules: Vec<(i32, i32)> = Vec::new();
    let mut total = 0;
    for line in file.trim().split("\n") {
        if line.contains("|") {
            rules.push(
                line.split("|")
                    .map(|e| e.parse::<i32>().unwrap_or_default())
                    .collect::<Vec<_>>()
                    .try_into()
                    .map(|arr: [i32; 2]| (arr[0], arr[1]))
                    .unwrap_or_default(),
            );

            continue;
        }

        if line == "" {
            continue;
        }

        let line_vec: Vec<i32> = line
            .split(",")
            .map(|e| e.parse::<i32>().unwrap_or_default())
            .collect();

        let mid = line_vec[line_vec.len() / 2];

        let update: HashMap<i32, usize> =
            line_vec
                .into_iter()
                .enumerate()
                .fold(HashMap::new(), |mut acc, ele| {
                    acc.insert(ele.1, ele.0);
                    acc
                });

        let mut is_right_order = true;
        for (left, right) in &rules {
            let contains_left = update.contains_key(left);
            let contains_right = update.contains_key(right);
            if contains_left && contains_right && update[&left] > update[&right] {
                is_right_order = false;
                break;
            }
        }

        if is_right_order {
            total += mid;
        }
    }
    total
}
