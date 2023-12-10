package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	tp       string
	distance int
	pos      []int
	child    *Node
	isInit   bool
	id       string
}

type Network struct {
	nodeMap [][]*Node
	grid    [][]rune
	start   *Node
}

func playPart0(fileName string) int {
	lines := readFile(fileName)
	fmt.Println(lines)
	return 0
}

func parseForPart0(input []string) *Network {

	rows := len(input)
	cols := len(input[0])
	ans := &Network{
		nodeMap: make([][]*Node, rows),
		grid:    make([][]rune, rows),
	}
	for row, line := range input {
		currChars := make([]rune, cols)
		currNodes := make([]*Node, cols)
		for col, char := range line {
			currChars[col] = char
			currNodes[col] = &Node{
				tp:     string(char),
				pos:    []int{row, col},
				isInit: char == 'S',
			}
		}
	}
}

func playPart1(fileName string) int {

	return 0
}

func main() {
	retVal := playPart0("test0.txt")
	fmt.Println(retVal)
	if retVal != 0 {
		panic("Test 0_0 failed")
	}
	fmt.Println("Test 0_0 passed")

	retVal = playPart0("test1.txt")
	fmt.Println(retVal)
	if retVal != 0 {
		panic("Test 0_1 failed")
	}
	fmt.Println("Test 0_1 passed")

	retVal = playPart0("input.txt")
	fmt.Println(retVal)
	if retVal != 0 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 0 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = playPart1("input.txt")
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
