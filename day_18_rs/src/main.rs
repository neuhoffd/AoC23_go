use core::fmt;
use std::{fs, string, vec};

#[derive(Debug)]
struct Command {
    cmd: char,
    val: i64,
}

#[derive(Debug, Clone)]
struct Point {
    x: i64,
    y: i64,
}

impl std::fmt::Display for Point {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        writeln!(f, "({},{})", self.x, self.y)
    }
}

fn get_direction(cmd: char) -> (i64, i64) {
    match cmd {
        'R' => (0, 1),
        'L' => (0, -1),
        'U' => (-1, 0),
        'D' => (1, 0),
        _ => {
            panic!("There should be a command here")
        }
    }
}

fn main() {
    let mut res: i64;

    assert_eq!(play(".\\test0.txt", true), 62);
    assert_eq!(play(".\\input.txt", true), 95356);

    assert_eq!(play(".\\test0.txt", false), 952408144115);
    assert_eq!(play(".\\input.txt", false), 92291468914147);
}

fn play(path: &str, part1: bool) -> i64 {
    let parsed: Vec<Point> = parse(
        &fs::read_to_string(path)
            .map(string::String::into_boxed_str)
            .unwrap(),
        part1,
    );

    let area = area(&parsed);
    println!(
        "Part {}: \nPath: {} \nArea: {}",
        if part1 { 1 } else { 2 },
        path,
        area
    );
    area
}

fn area(points: &Vec<Point>) -> i64 {
    let mut area: i64 = 0;
    let n: usize = points.len();
    let mut j: usize = n - 1;
    let mut perimeter: i64 = 0;

    for i in 0..n {
        area += (points[j].x + points[i].x) * (points[j].y - points[i].y);
        perimeter += (points[j].x - points[i].x).abs() + (points[j].y - points[i].y).abs();
        j = i;
    }
    area / 2 + perimeter / 2 + 1
}

fn parse(input: &str, part1: bool) -> Vec<Point> {
    let pos: Point = Point { x: 0, y: 0 };
    let cmds = if part1 {
        input
            .lines()
            .map(|line| -> Command {
                let mut splits = line.split_whitespace();
                Command {
                    cmd: splits.next().unwrap().chars().next().unwrap(),
                    val: splits.next().unwrap().parse().unwrap(),
                }
            })
            .collect::<Vec<Command>>()
    } else {
        input
            .lines()
            .map(|line| -> Command {
                let splits = line.split_whitespace();
                let hex = splits.skip(2).next().unwrap();
                Command {
                    val: i64::from_str_radix(&hex[2..7], 16).unwrap(),
                    cmd: match &hex[hex.len() - 2..hex.len() - 1] {
                        "0" => 'R',
                        "1" => 'D',
                        "2" => 'L',
                        "3" => 'U',
                        _ => panic!("This shouldn't happen"),
                    },
                }
            })
            .collect::<Vec<Command>>()
    };

    let mut ans: Vec<Point> = vec![Point { x: 0, y: 0 }];
    ans.append(
        &mut cmds
            .iter()
            .scan(pos, |pos: &mut Point, cmd: &Command| -> Option<Point> {
                let dir: (i64, i64);
                dir = get_direction(cmd.cmd);

                pos.x = pos.x + dir.0 * cmd.val;
                pos.y = pos.y + dir.1 * cmd.val;
                Some(pos.clone())
            })
            .collect::<Vec<Point>>(),
    );
    ans.split_last().unwrap().1.to_vec()
}
