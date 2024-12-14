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

fn move_direction(guard_position: &mut Position, lines: &mut Vec<Vec<char>>) {
    match &guard_position.direction {
        up_or_down @ (Direction::Up | Direction::Down) => {
            let cal_next_move = |cur_move: i32| -> i32 {
                match up_or_down {
                    Direction::Up => cur_move - 1,
                    _ => cur_move + 1,
                }
            };
            let guard_pos_row = guard_position.row as i32;
            let guard_pos_col = guard_position.col as usize;
            let mut next_move = cal_next_move(guard_pos_row);

            while next_move > 0 && lines[next_move as usize][guard_pos_col] != '#' {
                println!("next_move: {next_move}");

                // swap
                let temp = lines[guard_position.row as usize][guard_pos_col];
                lines[guard_position.row as usize][guard_pos_col] =
                    lines[next_move as usize][guard_pos_col];
                lines[next_move as usize][guard_pos_col] = temp;

                // update guard_position
                guard_position.row = next_move;

                // show current location
                println!("=========");
                for line in &mut *lines {
                    println!("{:?}", line);
                }
                println!("=========");

                next_move = cal_next_move(guard_position.row);
            }
        }
        left_or_right @ (Direction::Left | Direction::Right) => {
            let cal_next_move = |cur_move: i32| -> i32 {
                match left_or_right {
                    Direction::Left => cur_move - 1,
                    _ => cur_move + 1,
                }
            };
            let guard_pos_row = guard_position.row as usize;
            let guard_pos_col = guard_position.col as i32;
            let mut next_move = cal_next_move(guard_pos_col);
            while next_move < (lines[0].len() as i32)
                && lines[guard_pos_row][next_move as usize] != '#'
            {
                println!("next_move: {next_move}");

                let temp = lines[guard_pos_row][guard_pos_col as usize];
                lines[guard_pos_row][guard_pos_col as usize] =
                    lines[guard_pos_row][next_move as usize];
                lines[guard_pos_row][next_move as usize] = temp;

                // update guard position
                guard_position.col = next_move;

                // show current location
                println!("=========");
                for line in &mut *lines {
                    println!("{:?}", line);
                }
                println!("=========");

                next_move = cal_next_move(guard_position.col);
            }
        }
    }
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
    println!("{:?}", guard_position);

    while guard_position.col > 0 && guard_position.col < max_col as i32
        || guard_position.row > 0 && guard_position.row < max_row as i32
    {
        move_direction(&mut guard_position, &mut lines);
        guard_position.direction = change_direction(&guard_position.direction);
    }

    0
}
