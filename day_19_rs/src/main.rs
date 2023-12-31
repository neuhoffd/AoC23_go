use regex::Regex;
use std::collections::HashMap;
use std::fs;

#[derive(Debug)]
struct Part {
    x: isize,
    m: isize,
    a: isize,
    s: isize,
}

impl Part {
    fn parse(input: &str) -> Part {
        let trimmed = input.trim().trim_end_matches(['\r', '}']);
        let splits: Vec<&str> = trimmed[1..trimmed.len()].split(",").collect();
        let mut ans = Part {
            x: 0,
            m: 0,
            a: 0,
            s: 0,
        };
        for split in splits {
            let (target, val) = split.split_once("=").unwrap();
            let val: isize = val.parse().unwrap();
            ans.set(&target.chars().next().unwrap(), val);
        }
        ans
    }

    fn set(&mut self, target: &char, val: isize) {
        match target {
            'x' => self.x = val,
            'm' => self.m = val,
            'a' => self.a = val,
            's' => self.s = val,
            _ => panic!("Invalid Part Property in Set"),
        }
    }

    fn get(&self, target: &char) -> isize {
        match target {
            'x' => self.x,
            'm' => self.m,
            'a' => self.a,
            's' => self.s,
            _ => panic!("Invalid Part Property in Get"),
        }
    }

    fn rating(&self) -> isize {
        self.x + self.m + self.a + self.s
    }

    fn evaluate(&self, wfs: &HashMap<String, Workflow>) -> &Destination {
        let mut curr = wfs.get("in").unwrap();
        loop {
            match curr.eval(self) {
                Destination::Accept => {
                    println!("{:?} accepted", &self);
                    return &Destination::Accept;
                }
                Destination::Reject => {
                    println!("{:?} rejected", &self);
                    return &Destination::Reject;
                }
                Destination::Workflow(name) => curr = wfs.get(name).unwrap(),
            };
        }
    }
}

#[derive(Debug, Clone)]
struct Workflow {
    name: String,
    rules: Vec<Rule>,
}

impl Workflow {
    fn parse(input: &str) -> Workflow {
        let (name, rules) = input
            .trim()
            .trim_end_matches(['\r', '}'])
            .split_once("{")
            .unwrap();
        Workflow {
            name: name.to_string(),
            rules: rules
                .split(",")
                .into_iter()
                .map(|rule| Rule::parse(rule))
                .collect(),
        }
    }
    fn eval(&self, part: &Part) -> &Destination {
        self.rules.iter().find_map(|rule| rule.eval(part)).unwrap()
    }
}

#[derive(Debug, PartialEq, Clone, Eq)]
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

#[derive(Debug, Clone)]
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
            Rule::Fallthrough(Destination::parse(&input))
        }
    }

    fn eval(&self, part: &Part) -> Option<&Destination> {
        match self {
            Rule::Eval(cond, dest) => {
                if cond.evaluate(part) {
                    Some(dest)
                } else {
                    None
                }
            }
            Rule::Fallthrough(dest) => Some(dest),
        }
    }
}

#[derive(Debug, Clone, Copy)]
enum Condition {
    Greater(char, isize),
    Smaller(char, isize),
}

impl Condition {
    fn evaluate(&self, part: &Part) -> bool {
        match self {
            Condition::Greater(target, val) => part.get(target) > *val,
            Condition::Smaller(target, val) => part.get(target) < *val,
        }
    }

    fn invert(&self) -> Condition {
        match self {
            Condition::Greater(target, val) => Condition::Smaller(*target, *val + 1),
            Condition::Smaller(target, val) => Condition::Greater(*target, *val - 1),
        }
    }
}

fn determine_paths(
    wfs: &HashMap<String, Workflow>,
    curr: &str,
    path_history: &Vec<Condition>,
) -> Vec<Vec<Condition>> {
    let mut paths: Vec<Vec<Condition>> = vec![];
    let curr_wfs = wfs.get(curr).unwrap();
    let mut conditions_encountered = vec![];

    for rule in &curr_wfs.rules {
        let mut new_conditions = path_history.clone();
        new_conditions.extend(conditions_encountered.clone());
        match rule {
            Rule::Eval(cond, dest) => {
                new_conditions.push(*cond);
                conditions_encountered.push(cond.invert());
                match dest {
                    Destination::Accept => paths.push(new_conditions),
                    Destination::Reject => {}
                    Destination::Workflow(wf_name) => {
                        paths.extend(determine_paths(wfs, wf_name.as_str(), &new_conditions))
                    }
                }
            }
            Rule::Fallthrough(dest) => match dest {
                Destination::Accept => paths.push(new_conditions),
                Destination::Reject => {}
                Destination::Workflow(wf_name) => {
                    paths.extend(determine_paths(wfs, wf_name.as_str(), &new_conditions))
                }
            },
        }
    }
    paths
}

fn compute_combinations(path: &Vec<Condition>) -> isize {
    let mut min_part = Part {
        x: 1,
        m: 1,
        a: 1,
        s: 1,
    };
    let mut max_part = Part {
        x: 4000,
        m: 4000,
        a: 4000,
        s: 4000,
    };
    for cond in path {
        match cond {
            Condition::Greater(target, val) => {
                if *val > min_part.get(target) {
                    min_part.set(target, *val + 1);
                }
            }
            Condition::Smaller(target, val) => {
                if *val < max_part.get(target) {
                    max_part.set(target, *val - 1);
                }
            }
        }
    }
    (max_part.x - min_part.x + 1)
        * (max_part.m - min_part.m + 1)
        * (max_part.a - min_part.a + 1)
        * (max_part.s - min_part.s + 1)
}

fn main() {
    assert_eq!(play(".\\test0.txt", true), 19114);
    assert_eq!(play(".\\input.txt", true), 489392);

    assert_eq!(play(".\\test0.txt", false), 167409079868000);
    assert_eq!(play(".\\input.txt", false), 134370637448305);
}

fn play(path: &str, part1: bool) -> isize {
    let (wfs, pts) = parse(&fs::read_to_string(path).unwrap());
    if part1 {
        let res: isize = pts
            .iter()
            .filter(|part| *part.evaluate(&wfs) == Destination::Accept)
            .map(|part| part.rating())
            .sum();
        println!("Part1: {}, Result {}", part1, res);
        res
    } else {
        let paths = determine_paths(&wfs, "in", &vec![]);
        let res = paths
            .into_iter()
            .map(|path| compute_combinations(&path))
            .sum::<isize>();
        println!("Part1: {}, Result {}", part1, res);
        res
    }
}

fn parse_workflows(workflows: &str) -> HashMap<String, Workflow> {
    workflows
        .split("\n")
        .into_iter()
        .map(|line| Workflow::parse(line))
        .map(|wf| (wf.name.clone(), wf))
        .collect()
}

fn parse_parts(input: &str) -> Vec<Part> {
    input
        .split("\n")
        .into_iter()
        .map(|line| Part::parse(line))
        .collect()
}

// Vec<Part>, HashMap<String, Workflow>
fn parse(input: &String) -> (HashMap<String, Workflow>, Vec<Part>) {
    let (workflows, parts) = input.split_once("\r\n\r\n").unwrap();
    let wfs = parse_workflows(workflows);
    let pts = parse_parts(parts);
    (wfs, pts)
}
