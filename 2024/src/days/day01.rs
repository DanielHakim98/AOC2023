use std::{collections::HashMap, fs::read_to_string};

use crate::{Solution, SolutionPair};

///////////////////////////////////////////////////////////////////////////////

pub fn solve() -> SolutionPair {
    let file = read_to_string("inputs/day01.txt").expect("Failed to open input file");
    let lines: Vec<&str> = file.trim().split("\n").collect();

    let len = lines.len();
    let mut left_vec: Vec<i32> = Vec::with_capacity(len);
    let mut right_vec: Vec<i32> = Vec::with_capacity(len);

    let mut occurences: HashMap<i32, i32> = HashMap::new();

    for line in lines {
        let v: Vec<i32> = line
            .split_whitespace()
            .map(|e| e.parse::<i32>().unwrap())
            .collect();
        left_vec.push(v[0]);
        right_vec.push(v[1]);

        *occurences.entry(v[1]).or_insert(0) += 1;
    }

    merge_sort(&mut left_vec);
    merge_sort(&mut right_vec);

    let mut total = 0;
    let mut total_2 = 0;
    for i in 0..len {
        let left_val = left_vec[i];
        let right_val = right_vec[i];

        total += {
            let diff = left_val - right_val;
            if diff < 0 {
                -diff
            } else {
                diff
            }
        };

        if occurences.contains_key(&left_val) {
            total_2 += left_val * occurences[&left_val];
        }
    }

    let sol1: u64 = total as u64;
    let sol2: u64 = total_2 as u64;

    (Solution::from(sol1), Solution::from(sol2))
}

fn merge_sort(vec: &mut [i32]) {
    let len = vec.len();
    if len <= 1 {
        return;
    }

    let mut temp = vec.to_vec();

    let mid = len / 2;
    let (left, right) = vec.split_at_mut(mid);

    merge_sort(left);
    merge_sort(right);
    merge(left, right, &mut temp);

    vec.copy_from_slice(&temp);
}

fn merge(left: &[i32], right: &[i32], arr: &mut [i32]) {
    let (mut i, mut j, mut k) = (0, 0, 0);
    while i < left.len() && j < right.len() {
        if left[i] < right[j] {
            arr[k] = left[i];
            i += 1;
        } else {
            arr[k] = right[j];
            j += 1;
        }
        k += 1;
    }

    while i < left.len() {
        arr[k] = left[i];
        i += 1;
        k += 1;
    }

    while j < right.len() {
        arr[k] = right[j];
        j += 1;
        k += 1;
    }
}
