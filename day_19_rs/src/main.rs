use regex::Regex;
use std::collections::HashMap;
use std::fs;

#[derive(Debug)]
struct Workflow {
    name: String,
    rules: Vec<Rule>,
}

impl Workflow {
    fn parse(input: &str) -> Workflow {
        let (name, rules) = input[0..input.len() - 1].trim().split_once("{").unwrap();
        Workflow {
            name: name.to_string(),
            rules: rules
                .split(",")
                .into_iter()
                .map(|rule| Rule::parse(rule))
                .collect(),
        }
    }
}

#[derive(Debug)]
struct Part {
    x: isize,
    m: isize,
    a: isize,
    s: isize,
}

#[derive(Debug)]
enum Destination {
    Accept,
    Reject,
    Workflow(String),
}

impl Destination {
    fn parse(input: &str) -> Destination {
        match input {
            "A" => Destination::Accept,
            "R" => Destination::Reject,
            _ => Destination::Workflow(input.to_string()),
        }
    }
}

#[derive(Debug)]
enum Rule {
    Eval(Condition, Destination),
    Fallthrough(Destination),
}

impl Rule {
    fn parse(input: &str) -> Rule {
        let re =
            Regex::new(r"(?<cat>[xmas])(?<comp>[<>])(?<val>[0-9]+):(?<target>[a-zA-Z]+)").unwrap();
        if input.contains(['>', '<']) {
            let Some(captures): Option<regex::Captures<'_>> = re.captures(input) else {
                panic!("RegexCaptureFault")
            };
            match captures["comp"].chars().next().unwrap() {
                '>' => Rule::Eval(
                    Condition::Greater(
                        captures["cat"].chars().next().unwrap(),
                        captures["val"].parse().unwrap(),
                    ),
                    Destination::parse(&captures["target"]),
                ),
                '<' => Rule::Eval(
                    Condition::Smaller(
                        captures["cat"].chars().next().unwrap(),
                        captures["val"].parse().unwrap(),
                    ),
                    Destination::parse(&captures["target"]),
                ),
                _ => panic!("Rule Parse Error"),
            }
        } else {
            Rule::Fallthrough(Destination::parse(input))
        }
    }
}

#[derive(Debug)]
enum Condition {
    Greater(char, isize),
    Smaller(char, isize),
}

fn main() {
    assert_eq!(play(".\\test0.txt", true), 19114);
    assert_eq!(play(".\\input.txt", true), -1);

    assert_eq!(play(".\\test0.txt", false), -1);
    assert_eq!(play(".\\input.txt", false), -1);
}

fn play(path: &str, part1: bool) -> isize {
    let parsed: () = parse(&fs::read_to_string(path).unwrap(), part1);
    -1
}

fn parse_workflows(workflows: &str) -> HashMap<String, Workflow> {
    workflows
        .split("\n")
        .into_iter()
        .map(|line| Workflow::parse(line))
        .map(|wf| (wf.name.clone(), wf))
        .collect()
    /*for line in workflows.split("\n") {
        let (name, mut rules) = line.split_once("{").unwrap();
        let mut wf = Workflow {
            name: name.to_string(),
            rules: vec![],
        };
        rules = &rules[0..rules.len() - 1];
        for rule in rules.split(",") {
            if rule.contains(['>', '<']) {
                let Some(captures): Option<regex::Captures<'_>> = re.captures(rule) else {
                    panic!("RegexCaptureFault")
                };

                wf.rules.push(Destination {
                    greater: captures["comp"].chars().next().unwrap(),
                    val: captures["val"].parse::<isize>().unwrap(),
                    target: captures["target"].to_string(),
                })
            } else {
                wf.rules.push(Destination {
                    greater: ' ',
                    val: -1,
                    target: rule.to_string(),
                })
            }
        }
        ans.insert(wf.name.clone(), wf);
        dbg!(&ans);
    }
    ans*/
}

fn parse_parts(input: &str) -> Vec<Part> {
    let mut ans: Vec<Part> = vec![];
    /*input.split("\n").into_iter().map(|item| -> () {
        ans.push(Part {
            map: HashMap::from([]),
        });
    });*/
    ans
}

// Vec<Part>, HashMap<String, Workflow>
fn parse(input: &String, part1: bool) -> () {
    if let (workflows, parts) = input.split_once("\r\n\r\n").unwrap() {
        let wfs = Workflow::parse(workflows);
        dbg!(wfs);
        let pts = parse_parts(parts);
    } else {
        panic!("Parsing Error")
    }
}
