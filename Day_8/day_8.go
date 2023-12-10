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
	nodeMap      map[string]*Node
	currentPos   *Node
	instructions string
	steps        int
}

func (network *Network) makeNodeIfNotExists(id string) *Node {
	newNode, ok := network.nodeMap[id]
	if !ok {
		newNode = &Node{
			id:       id,
			isRoot:   id == "AAA",
			isTarget: id == "ZZZ",
		}
		network.nodeMap[id] = newNode
	}
	return newNode
}

func (network *Network) traverse() {
	stepCounter := 0
	for !network.currentPos.isTarget {
		currInstruction := stepCounter % len(network.instructions)
		stepCounter++
		if network.instructions[currInstruction] == 'L' {
			network.currentPos = network.currentPos.children[0]
		} else {
			network.currentPos = network.currentPos.children[1]
		}
	}
	network.steps = stepCounter
}

func playPart1(fileName string) int {
	lines := readFile(fileName)
	fmt.Println(lines)
	network := parseForPart1(lines)
	fmt.Println(network)
	network.traverse()
	return network.steps
}

func parseForPart1(input []string) Network {
	ans := Network{instructions: strings.TrimSpace(input[0]), nodeMap: make(map[string]*Node)}
	r := regexp.MustCompile(`\w+`)
	for _, line := range input[2:] {
		nodeStrings := r.FindAllString(line, 3)
		fmt.Println(nodeStrings)
		currNodeId := nodeStrings[0]
		currNode := ans.makeNodeIfNotExists(currNodeId)
		if currNode.isRoot {
			ans.currentPos = currNode
		}
		for i := 1; i <= len(currNode.children); i++ {
			currChild := ans.makeNodeIfNotExists(nodeStrings[i])
			currNode.children[i-1] = currChild
		}
	}
	return ans
}

func playPart2(fileName string) int {

	return 0
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
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != 16343 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart2("test0.txt")
	fmt.Println(retVal)
	if retVal != 0 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = playPart2("input.txt")
	fmt.Println(retVal)
	if retVal != 0 {
		panic("Part 1 failed")
	}
	fmt.Println("Part 1 passed")
}

func readFile(fileName string) []string {
	bytes, _ := os.ReadFile(fileName)
	input := string(bytes)
	splitInput := strings.Split(input, "\n")
	return splitInput
}
