package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Node struct {
	tp       string
	distance int
	pos      []int
	children []*Node
}

type Network struct {
	nodeMap    [][]*Node
	start      *Node
	rows, cols int
}

func (network *Network) print() {
	for _, nodes := range network.nodeMap {
		for _, node := range nodes {
			fmt.Printf("%s", node.tp)
		}
		fmt.Printf("\n")
	}
}

func (network *Network) countEnclosed() int {
	ans := 0
	for row := 0; row < network.rows; row++ {
		verticalFound := 0
		last := "-"
		for col := 0; col < network.cols; col++ {
			currNode := network.nodeMap[row][col]
			if currNode.distance < 0 {
				currNode.tp = "."
			}
			switch currNode.tp {
			case "|":
				{
					verticalFound++
				}
			case "-":
				{
					continue
				}
			case ".":
				{
					if verticalFound%2 != 0 {
						ans++
						currNode.tp = "I"
					} else {
						currNode.tp = " "
					}
				}

			case "F", "L", "J", "7":
				{
					if last == "-" {
						last = currNode.tp
						continue
					}
					if (last == "L" && currNode.tp == "7") || (last == "F" && currNode.tp == "J") {
						verticalFound++
					}
					last = "-"
				}
			}
		}
	}
	network.print()

	return ans
}

func playPart0(fileName string) int {
	lines := readFile(fileName)
	network := parseForPart0(lines)
	network.init()
	result := network.traverseFromStart()
	return result
}

func playPart1(fileName string) int {
	lines := readFile(fileName)
	network := parseForPart0(lines)
	network.init()
	network.traverseFromStart()
	result := network.countEnclosed()
	return result
}

func (network *Network) traverseFromStart() int {
	furthestFound := false
	currNodes := make([]*Node, len(network.start.children))
	numPaths := len(currNodes)
	for i := 0; i < numPaths; i++ {
		currNodes[i] = network.start.children[i]
		currNodes[i].distance = 1
	}
	for !furthestFound {
		oldDistance := currNodes[0].distance
		for i := 0; i < numPaths; i++ {
			for _, child := range currNodes[i].children {
				if child.distance < 0 {
					currNodes[i] = child
					break
				}
			}
		}
		furthestFound = true
		for i := 0; i < numPaths; i++ {
			currNodes[i].distance = oldDistance + 1
			furthestFound = currNodes[i] == currNodes[0] && furthestFound
		}
	}
	return currNodes[0].distance
}

func (network *Network) init() {
	offsets := map[string][][]int{
		"|": {{-1, 0}, {1, 0}},
		"-": {{0, -1}, {0, 1}},
		"L": {{-1, 0}, {0, 1}},
		"J": {{-1, 0}, {0, -1}},
		"7": {{1, 0}, {0, -1}},
		"F": {{1, 0}, {0, 1}},
	}

	for _, currRow := range network.nodeMap {
		for _, currNode := range currRow {
			currPos := currNode.pos
			for _, offset := range offsets[currNode.tp] {
				candidatePos := []int{currPos[0] + offset[0], currPos[1] + offset[1]}
				if slices.Min(candidatePos) < 0 || candidatePos[0] >= network.rows || candidatePos[1] >= network.cols {
					continue
				}
				currNode.children = append(currNode.children, network.nodeMap[candidatePos[0]][candidatePos[1]])
			}
		}
	}
	startChildOffsets := make([][]int, 0)
	startNode := network.start
	for rowOffset := -1; rowOffset <= 1; rowOffset++ {
		for colOffset := -1; colOffset <= 1; colOffset++ {

			candidatePos := []int{startNode.pos[0] + rowOffset, startNode.pos[1] + colOffset}
			if slices.Min(candidatePos) < 0 || candidatePos[0] >= network.rows || candidatePos[1] >= network.cols || (colOffset == 0 && rowOffset == 0) {
				continue
			}
			candidateNode := network.nodeMap[startNode.pos[0]+rowOffset][startNode.pos[1]+colOffset]
			if slices.Contains(candidateNode.children, startNode) {
				startNode.children = append(startNode.children, candidateNode)
				startChildOffsets = append(startChildOffsets, []int{rowOffset, colOffset})
			}
		}
	}

	for k, v := range offsets {
		found := ((slices.Equal(v[0], startChildOffsets[0]) && slices.Equal(v[1], startChildOffsets[1])) ||
			(slices.Equal(v[1], startChildOffsets[0]) && slices.Equal(v[0], startChildOffsets[1])))
		if found {
			startNode.tp = k
			break
		}
	}
	network.start.distance = 0
}

func parseForPart0(input []string) *Network {
	rows := len(input)
	cols := len(strings.TrimSpace(input[0]))
	ans := &Network{
		nodeMap: make([][]*Node, rows),
		rows:    rows,
		cols:    cols,
	}
	for row, line := range input {
		line = strings.TrimSpace(line)
		currChars := make([]rune, cols)
		currNodes := make([]*Node, cols)
		for col, char := range line {
			currChars[col] = char
			currNodes[col] = &Node{
				tp:       string(char),
				pos:      []int{row, col},
				distance: -1,
			}
			if char == 'S' {
				ans.start = currNodes[col]
				continue
			}
		}
		ans.nodeMap[row] = currNodes
	}
	return ans
}

func main() {
	retVal := playPart0("test0.txt")
	fmt.Println(retVal)
	if retVal != 4 {
		panic("Test 0_0 failed")
	}
	fmt.Println("Test 0_0 passed")

	retVal = playPart0("test1.txt")
	fmt.Println(retVal)
	if retVal != 8 {
		panic("Test 0_1 failed")
	}
	fmt.Println("Test 0_1 passed")

	retVal = playPart0("input.txt")
	fmt.Println(retVal)
	if retVal != 6754 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart1("test2.txt")
	fmt.Println(retVal)
	if retVal != 4 {
		panic("Test 1_1 failed")
	}
	fmt.Println("Test 1_1 passed")

	retVal = playPart1("test3.txt")
	fmt.Println(retVal)
	if retVal != 8 {
		panic("Test 1_2 failed")
	}
	fmt.Println("Test 1_2 passed")

	retVal = playPart1("test4.txt")
	fmt.Println(retVal)
	if retVal != 10 {
		panic("Test 1_3 failed")
	}
	fmt.Println("Test 1_3 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != 567 {
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
