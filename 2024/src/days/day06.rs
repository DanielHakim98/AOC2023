use crate::{Solution, SolutionPair};
use std::fs::read_to_string;

pub fn solve() -> SolutionPair {
    let file = read_to_string("inputs/day06_test.txt").expect("Failed to open input file");
    let sol1: u64 = part1(&file) as u64;
    let sol2: u64 = 0 as u64;

    (Solution::from(sol1), Solution::from(sol2))
}

#[derive(Debug)]
struct Position {
    row: usize,
    col: usize,
}

enum Direction {
    Up,
    Right,
    Down,
    Left,
}

fn part1(content: &str) -> usize {
    let mut lines: Vec<Vec<char>> = content
        .trim()
        .split("\n")
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect();

    let max_row = lines.len();
    let max_col = lines[0].len();
    let mut guard_position = Position {
        row: max_row,
        col: max_col,
    };

    println!("{:?}", guard_position);

    for i in 0..lines.len() {
        let row_index = i;
        let col_index = lines[i]
            .iter()
            .position(|&r| r == '^')
            .unwrap_or_else(|| max_col);
        if col_index != max_col {
            guard_position.col = col_index;
            guard_position.row = row_index;
            break;
        }
    }

    let mut next_move = guard_position.row - 1;
    while next_move > 0 && lines[next_move][guard_position.col] != '#' {
        println!("next_move: {next_move}");
        let temp = lines[guard_position.row][guard_position.col];
        lines[guard_position.row][guard_position.col] = lines[next_move][guard_position.col];
        lines[next_move][guard_position.col] = temp;
        guard_position.row = next_move;
        for line in &lines {
            println!("{:?}", line);
        }
        next_move -= 1;
    }

    let mut next_move = guard_position.col + 1;
    while next_move < max_col && lines[guard_position.row][next_move] != '#' {
        println!("next_move: {next_move}");
        let temp = '>';
        lines[guard_position.row][guard_position.col] = lines[guard_position.row][next_move];
        lines[guard_position.row][next_move] = temp;
        guard_position.col = next_move;
        for line in &lines {
            println!("{:?}", line);
        }
        next_move += 1;
    }

    0
}
