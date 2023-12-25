use core::fmt;
use std::{fs, string, vec, iter::once, num};

#[derive(Debug)]
struct Command {
    cmd: char,
    val: i32,
}

#[derive(Debug)]
#[derive(Clone)]
struct Point {
    x: i32, y: i32,
}

impl std::fmt::Display for Point {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        writeln!(f,"({},{})", self.x, self.y)
    }
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
    let input = fs::read_to_string("C:\\workspace\\AoC23_go\\day_18_rs\\src\\input.txt")
        .map(string::String::into_boxed_str)
        .unwrap();

    println!("{}", input);

    let parsed: Vec<Point> = parse(&input);

    println!("{:?}", parsed);

    print_points(&parsed);
    println!("{}", calc_area(&parsed));
}

fn calc_area(points: &Vec<Point>) -> i32 {
    let mut area :i32= 0;
    let n: usize = points.len();
    let mut j: usize = n-1;
    let mut perimeter: i32 = 0;

    for i in 0..n {
        let term1 = points[j].x + points[i].x;
        let term2 = points[j].y - points[i].y;
        println!("Area {}, i {}, j {}, t1 {}, t2 {}, p {}", area / 2, i, j, term1, term2, perimeter);
        
        area += (points[j].x + points[i].x) * (points[j].y - points[i].y);       
        perimeter += (points[j].x - points[i].x).abs() + term2.abs();
        j = i; 
    }
    area / 2 + perimeter / 2 + 1
}

fn print_points(points: &Vec<Point>) {
    println!("{}",points.iter().map(|p| p.to_string()).collect::<Vec<String>>().join(""))
}

fn parse(input: &str) -> Vec<Point> {
    let mut pos : Point = Point { x: 0, y: 0 };

    let cmds =     
        input
        .lines()
        .map(|line| -> Command {
            let mut splits = line.split_whitespace();
            Command {
                cmd: splits.next().unwrap().chars().next().unwrap(),
                val: splits.next().unwrap().parse().unwrap(),
            }
        })
        .collect::<Vec<Command>>();

    let mut ans : Vec<Point> = vec![Point{x:0, y:0}];
    ans.append(&mut cmds.iter()
    .scan(pos,  |mut pos: &mut Point, cmd: &Command| -> Option<Point> {
        let dir: (i32, i32);
        dir = get_direction(cmd.cmd);

        pos.x = pos.x + dir.0 * cmd.val;
        pos.y = pos.y + dir.1 * cmd.val;
        Some(pos.clone())
    }).collect::<Vec<Point>>());
    ans.split_last().unwrap().1.to_vec()  
}

  /*
    )
    let points = firstPoint.into_iter()
        .chain(
        cmds
        .iter()
        .scan(pos,  |mut pos: &mut Point, cmd: &Command| -> Option<Point> {
            let dir: (i32, i32);
            dir = get_direction(cmd.cmd);
    
            pos.x = pos.x + dir.0 * cmd.val;
            pos.y = pos.y + dir.1 * cmd.val;
            Some(pos.clone())
        }).collect::<Vec<Point>>().iter()).collect();*/


    /*parsed.iter().fold(pos, |mut acc, cmd: &Command| {
        let dir: (i32, i32);
        dir = get_direction(cmd.cmd);

        acc = (acc.0 + dir.0 * cmd.val, acc.1 + dir.1 * cmd.val);
        let (row, col) = acc;
        min_max[0] = std::cmp::min(row, min_max[0]);
        min_max[1] = std::cmp::max(row, min_max[1]);
        min_max[2] = std::cmp::min(col, min_max[2]);
        min_max[3] = std::cmp::max(col, min_max[3]);
        acc
    });

    let max_row: usize = usize::try_from(min_max[1] - min_max[0]).unwrap();
    let max_col: usize = usize::try_from(min_max[3] - min_max[2]).unwrap();
    let mut grid: Vec<Vec<char>> = vec![vec!['.'; max_col + 1]; max_row + 1];
    let mut shifted_pos: (usize, usize) = (min_max[0] as usize, min_max[2] as usize);
    grid[shifted_pos.0][shifted_pos.1] = '#';*/

    /*for cmd in parsed.iter() {
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

    print_marked(grid)*/
/*
    
fn print_marked(grid: Vec<Vec<char>>) {
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
} */
