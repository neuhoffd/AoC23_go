package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	id               string
	isRoot, isTarget bool
	children         [2]*Node
}

type Network struct {
	nodeMap       map[string]*Node
	currPositions []*Node
	instructions  string
	steps         int
}

func (network *Network) makeNodeIfNotExists(id string, part2 bool) *Node {
	newNode, ok := network.nodeMap[id]
	if !ok {
		newNode = &Node{
			id:       id,
			isRoot:   id == "AAA" || (id[2] == 'A' && part2),
			isTarget: id == "ZZZ" || (id[2] == 'Z' && part2),
		}
		network.nodeMap[id] = newNode
	}
	return newNode
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers ...int) int {
	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func (network *Network) traversePart1() {
	stepCounter := 0
	allArrived := false
	for !allArrived {
		currInstruction := stepCounter % len(network.instructions)
		stepCounter++
		if stepCounter%100000 == 0 {
			fmt.Printf("Step:  %d, Positions %+v \n", stepCounter, network.currPositions)
		}
		allArrived = true
		for i := 0; i < len(network.currPositions); i++ {
			if !network.currPositions[i].isTarget {
				allArrived = false
			}
			if network.instructions[currInstruction] == 'L' {
				network.currPositions[i] = network.currPositions[i].children[0]
			} else {
				network.currPositions[i] = network.currPositions[i].children[1]
			}
		}
	}
	network.steps = stepCounter - 1
}

func (network *Network) traversePart2() {
	cycles := []int{}
	for i := 0; i < len(network.currPositions); i++ {
		stepCounter := 0
		fmt.Println("Traversing Position ", i)
		for {
			currInstruction := stepCounter % len(network.instructions)
			stepCounter++
			if stepCounter%100000 == 0 {
				fmt.Printf("Step:  %d, Positions %+v \n", stepCounter, network.currPositions)
			}
			if network.currPositions[i].isTarget {
				break
			}

			if network.instructions[currInstruction] == 'L' {
				network.currPositions[i] = network.currPositions[i].children[0]
			} else {
				network.currPositions[i] = network.currPositions[i].children[1]
			}
		}
		cycles = append(cycles, stepCounter-1)
	}
	network.steps = LCM(cycles...)
}

func playPart1(fileName string) int {
	lines := readFile(fileName)
	network := parse(lines, false)
	network.traversePart1()
	return network.steps
}

func playPart2(fileName string) int {
	lines := readFile(fileName)
	network := parse(lines, true)
	network.traversePart2()
	return network.steps
}

func parse(input []string, partTwo bool) Network {
	ans := Network{instructions: strings.TrimSpace(input[0]), nodeMap: make(map[string]*Node)}
	r := regexp.MustCompile(`\w+`)
	for _, line := range input[2:] {
		nodeStrings := r.FindAllString(line, 3)
		currNodeId := nodeStrings[0]
		currNode := ans.makeNodeIfNotExists(currNodeId, partTwo)
		if currNode.isRoot {
			ans.currPositions = append(ans.currPositions, currNode)
		}
		for i := 1; i <= len(currNode.children); i++ {
			currChild := ans.makeNodeIfNotExists(nodeStrings[i], partTwo)
			currNode.children[i-1] = currChild
		}
	}
	return ans
}

func main() {
	retVal := playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 2 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart1("test1.txt")
	fmt.Println(retVal)
	if retVal != 6 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != 16343 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart2("test2.txt")
	fmt.Println(retVal)
	if retVal != 6 {
		panic("Test 2 failed")
	}
	fmt.Println("Test 2 passed")

	retVal = playPart2("input.txt")
	fmt.Println(retVal)
	if retVal != 15299095336639 {
		panic("Part 3 failed")
	}
	fmt.Println("Part 3 passed")
}

func readFile(fileName string) []string {
	bytes, _ := os.ReadFile(fileName)
	input := string(bytes)
	splitInput := strings.Split(input, "\n")
	return splitInput
}
