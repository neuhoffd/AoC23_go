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