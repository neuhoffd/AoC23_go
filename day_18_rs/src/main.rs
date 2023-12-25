use std::{fs, string};

#[derive(Debug)]
struct Workflow {}

struct Rule {}

#[repr(u8)]
enum Categories {
    X = b'x',
    M = b'm',
    A = b'a',
    S = b's',
}

fn main() {
    assert_eq!(play(".\\test0.txt", true), 19114);
    assert_eq!(play(".\\input.txt", true), -1);

    assert_eq!(play(".\\test0.txt", false), -1);
    assert_eq!(play(".\\input.txt", false), -1);
}

fn play(path: &str, part1: bool) -> i64 {
    let parsed: () = parse(&fs::read_to_string(path).unwrap(), part1);
    -1
}

fn parse_workflows(workflows: &str) {}

fn parse(input: &String, part1: bool) -> () {
    let (workflows, parts) = input.split_once("\n\n").unwrap();
}
