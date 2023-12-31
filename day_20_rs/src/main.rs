use regex::Regex;
use std::collections::HashMap;
use std::fs;

type Modules = HashMap<String, Module>;

#[derive(Debug)]
enum Module {
    FlipFlop(bool, Vec<String>),
    Conjunction(HashMap<String, Pulse>, Vec<String>),
    Broadcast(Vec<String>),
    Output(isize),
}

impl Module {
    fn parse(input: &String) -> (String, Module) {
        let (input, targets) = input.trim().split_once(" -> ").unwrap();
        if input.trim().starts_with("broadcast") {
            (
                String::from("broadcast"),
                Module::Broadcast(targets.split(", ").map(|s| s.trim().to_string()).collect()),
            )
        } else {
            match input.trim().chars().next().unwrap() {
                '&' => (
                    input[1..input.len()].to_string(),
                    Module::FlipFlop(
                        false,
                        targets.split(", ").map(|s| s.trim().to_string()).collect(),
                    ),
                ),
                '%' => (
                    input[1..input.len()].to_string(),
                    Module::Conjunction(
                        HashMap::new(),
                        targets.split(", ").map(|s| s.trim().to_string()).collect(),
                    ),
                ),
                _ => panic!("Shouldn't be here"),
            }
        }
    }
}

#[derive(Debug)]
enum Pulse {
    High(Module, Module),
    Low(Module, Module),
}

fn main() {
    assert_eq!(play(".\\test0.txt", true), 32000000);
    assert_eq!(play(".\\test1.txt", true), 11687500);
    assert_eq!(play(".\\input.txt", true), 0);

    assert_eq!(play(".\\test0.txt", false), 0);
    assert_eq!(play(".\\input.txt", false), 0);
}

fn play(path: &str, part1: bool) -> isize {
    let input = parse(&fs::read_to_string(path).unwrap());
    dbg!(&input);
    if part1 {
        -1
    } else {
        -1
    }
}

// Vec<Part>, HashMap<String, Workflow>
fn parse(input: &String) -> Modules {
    input
        .lines()
        .map(|line| line.trim())
        .map(|line| Module::parse(&line.to_string()))
        .collect()
}
