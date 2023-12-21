use std::{fs, string, vec};

#[derive(Debug)]
struct Command {
    cmd: char,
    val: i32,
}

fn get_direction(cmd: char) -> (i32, i32) {
    let dir: (i32, i32);
    match cmd {
        'R' => dir = (0, 1),
        'L' => dir = (0, -1),
        'U' => dir = (-1, 0),
        'D' => dir = (1, 0),
        _ => {
            dir = (0, 0);
            println!("There should be a command here")
        }
    }
    dir
}

fn main() {
    let input = fs::read_to_string("F:\\workspace\\AoC23_go\\day_18_rs\\src\\test0.txt")
        .map(string::String::into_boxed_str)
        .unwrap();

    println!("{}", input);

    let parsed = parse(&input);

    println!("{:?}", parsed);

    let pos: (i32, i32) = (0, 0);
    let mut min_max = [0i32; 4];

    parsed.iter().fold(pos, |mut acc, cmd: &Command| {
        let dir: (i32, i32);
        dir = get_direction(cmd.cmd);

        acc = (acc.0 + dir.0 * cmd.val, acc.1 + dir.1 * cmd.val);
        let (row, col) = acc;
        if row > min_max[1] {
            min_max[1] = row;
        }
        if row < min_max[0] {
            min_max[0] = row;
        }
        if col > min_max[3] {
            min_max[3] = col;
        }
        if col < min_max[2] {
            min_max[2] = col;
        }
        acc
    });

    let max_row: usize = usize::try_from(min_max[1] - min_max[0]).unwrap();
    let max_col: usize = usize::try_from(min_max[3] - min_max[2]).unwrap();
    let mut grid: Vec<Vec<char>> = vec![vec!['.'; max_col + 1]; max_row + 1];
    let mut shifted_pos: (usize, usize) = (min_max[0] as usize, min_max[2] as usize);
    grid[shifted_pos.0][shifted_pos.1] = '#';

    for cmd in parsed.iter() {
        let dir: (i32, i32) = get_direction(cmd.cmd);
        for _i in 0..cmd.val {
            shifted_pos = (
                ((shifted_pos.0 as i32) + dir.0) as usize,
                ((shifted_pos.1 as i32) + dir.1) as usize,
            );
            let (row, col) = shifted_pos;
            grid[row][col] = '#';
        }
    }
    println!("{:?}", grid);

    printMarked(grid)
}

fn printMarked(grid: Vec<Vec<char>>) {
    let mut sum = 0;
    for row in grid.iter() {
        for ele in row.iter() {
            print!("{}", ele);
            if *ele == '#' {
                sum += 1
            }
        }
        print!("\n")
    }
    print!("{}", sum)
}
fn parse(input: &str) -> Vec<Command> {
    input
        .lines()
        .map(|line| -> Command {
            let mut splits = line.split_whitespace();
            Command {
                cmd: splits.next().unwrap().chars().next().unwrap(),
                val: splits.next().unwrap().parse().unwrap(),
            }
        })
        .collect()
}
