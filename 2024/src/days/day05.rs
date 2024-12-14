use crate::{Solution, SolutionPair};
use std::{collections::HashMap, fs::read_to_string};

pub fn solve() -> SolutionPair {
    let file = read_to_string("inputs/day05_basic.txt").expect("Failed to open input file");

    let sol1: u64 = part1(&file) as u64;
    let sol2: u64 = part2(&file) as u64;

    (Solution::from(sol1), Solution::from(sol2))
}

fn part2(file: &str) -> i32 {
    let mut rules: HashMap<usize, Vec<usize>> = HashMap::new();
    let mut total = 0;
    for line in file.trim().split("\n") {
        if line == "" {
            println!("{:#?}", rules);
            continue;
        }

        if line.contains("|") {
            let edges: Vec<i32> = line
                .split("|")
                .map(|e| e.parse::<i32>().unwrap_or_default())
                .collect();

            let (left, right) = (edges[0], edges[1]);
            rules
                .entry(left as usize)
                .or_insert_with(|| Vec::new())
                .push(right as usize);
            rules.entry(right as usize).or_insert_with(|| Vec::new());
            continue;
        }

        let update: Vec<i32> = line
            .split(",")
            .map(|e| e.parse::<i32>().unwrap_or_default())
            .collect();
        let mut in_order = true;
        for i in 0..update.len() - 1 {
            let node = update[i] as usize;
            let next_node = update[i + 1] as usize;
            let temp: Vec<usize> = Vec::new();
            let out_nodes = &rules.get(&(node)).unwrap_or(&temp);
            print!("node({})", node);
            print!("--->{:?} ", out_nodes);
            println!("next_node = {}", next_node);

            let is_next_node_in_outnodes = out_nodes.contains(&next_node);
            let is_next_node_as_dependecies = rules.contains_key(&next_node);
            if !is_next_node_in_outnodes && is_next_node_as_dependecies {
                let next_node_outnodes = &rules.get(&next_node).unwrap_or(&temp);
                let is_cur_node_in_next_node_outnodes = next_node_outnodes.contains(&node);
                if is_cur_node_in_next_node_outnodes {
                    print!("next_node({})", next_node);
                    println!("--->{:?} ", next_node_outnodes);
                    in_order = false;
                    break;
                }
            }
        }
        if !in_order {
            let mut sorted: Vec<usize> = update.clone().into_iter().map(|e| e as usize).collect();

            sorted.sort_by(|&node, &next_node| {
                let left_len = rules.get(&node).map_or(0, |out_nodes| out_nodes.len());
                let right_len = rules.get(&next_node).map_or(0, |out_nodes| out_nodes.len());
                right_len.cmp(&left_len)
            });
            total += sorted[&sorted.len() / 2];
        }

        println!()
    }

    total as i32
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
