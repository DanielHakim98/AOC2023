use crate::{Solution, SolutionPair};
use std::{collections::HashMap, fs::read_to_string};

pub fn solve() -> SolutionPair {
    let file = read_to_string("inputs/day06.txt").expect("Failed to open input file");
    let sol1: u64 = part1(&file) as u64;
    let sol2: u64 = 0 as u64;

    (Solution::from(sol1), Solution::from(sol2))
}

#[derive(Debug)]
struct Position {
    row: i32,
    col: i32,
    direction: Direction,
}

#[derive(Debug)]
enum Direction {
    Up,
    Right,
    Down,
    Left,
}

fn init_direction() -> Direction {
    Direction::Up
}

fn change_direction(cur_direction: &Direction) -> Direction {
    match cur_direction {
        Direction::Up => Direction::Right,
        Direction::Right => Direction::Down,
        Direction::Down => Direction::Left,
        Direction::Left => Direction::Up,
    }
}

fn move_direction(guard_position: &Position) -> (i32, i32) {
    match &guard_position.direction {
        Direction::Up => (-1, 0),
        Direction::Right => (0, 1),
        Direction::Down => (1, 0),
        Direction::Left => (0, -1),
    }
}

fn part1(content: &str) -> usize {
    let lines: Vec<Vec<char>> = content
        .trim()
        .split("\n")
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect();

    let max_row = lines.len();
    let max_col = lines[0].len();
    let mut guard_position = Position {
        row: max_row as i32,
        col: max_col as i32,
        direction: init_direction(),
    };

    for i in 0..lines.len() {
        let row_index = i;
        let col_index = lines[i]
            .iter()
            .position(|&r| r == '^')
            .unwrap_or_else(|| max_col);
        if col_index != max_col {
            guard_position.col = col_index as i32;
            guard_position.row = row_index as i32;
            break;
        }
    }

    let mut step_history: HashMap<(i32, i32), usize> = HashMap::new();
    while guard_position.row > 0
        && guard_position.row < (max_row - 1) as i32
        && guard_position.col > 0
        && guard_position.col < (max_col - 1) as i32
    {
        let mut next_row_pos = guard_position.row + move_direction(&guard_position).0;
        let mut next_col_pos = guard_position.col + move_direction(&guard_position).1;

        if lines[next_row_pos as usize][next_col_pos as usize] == '#' {
            guard_position.direction = change_direction(&guard_position.direction);
            next_row_pos = guard_position.row + move_direction(&guard_position).0;
            next_col_pos = guard_position.col + move_direction(&guard_position).1;
        }

        guard_position.row = next_row_pos;
        guard_position.col = next_col_pos;

        let coord = (guard_position.row, guard_position.col);
        *step_history.entry(coord).or_default() += 1;
    }
    // println!("{:#?}", step_history);
    step_history.len()
}
