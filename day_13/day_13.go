package main

import (
	"fmt"
	"os"
	"strings"
)

type Pattern struct {
	rows, cols []int
	nCols      int
}

func playPart0(fileName string) int {
	input := readFile(fileName)
	fmt.Println(input)
	patterns := parseForPart0(input)
	fmt.Println(patterns)
	return 0
}

func parseForPart0(input []string) []*Pattern {
	ans := make([]*Pattern, 0)
	currPattern := &Pattern{rows: make([]int, 0), cols: make([]int, 0)}
	ans = append(ans, currPattern)
	for _, line := range input {
		if len(line) == 0 {
			fmt.Println("New Pattern")
			ans = append(ans, currPattern)
			currPattern = &Pattern{rows: make([]int, 0), cols: make([]int, 0)}
			continue
		}
		currPattern.nCols = len(line)
		currRow := 0
		for _, char := range line {
			switch char {
			case '.':
				{
					currRow = currRow<<1 + 0
				}
			case '#':
				{
					currRow = currRow<<1 + 1
				}
			default:
				panic("shouldn't be here")
			}
		}
		fmt.Printf("Line: %s, Bits: %b \n", line, currRow)
		currPattern.rows = append(currPattern.rows, currRow)
	}
	for _, pattern := range ans {
		for i := 0; i < pattern.nCols; i++ {
			pattern.cols = append(pattern.cols, 0)
			for j := 0; j < len(pattern.rows); j++ {
				addVal := pattern.rows[j] & (1 << (pattern.nCols - 1 - i))
				if addVal != 0 {
					addVal = 1
				}
				pattern.cols[i] = pattern.cols[i]<<1 + addVal
			}
			fmt.Printf("Pattern: %+v\n Bits: %b \n", pattern, pattern.cols[i])
		}
	}
	return ans
}

func playPart1(fileName string) int {

	return 0
}

func main() {
	retVal := playPart0("test0.txt")
	fmt.Println(retVal)
	if retVal != 405 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

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
	for idx := 0; idx < len(splitInput); idx++ {
		splitInput[idx] = strings.TrimSpace(splitInput[idx])
	}
	return splitInput
}
