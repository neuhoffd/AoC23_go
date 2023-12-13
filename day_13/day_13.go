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

func findReflectionLine(axes []int) int {
	ans := 0
	for i := 0; i < len(axes)-1; i++ {
		maxExtent := i + 1
		if maxExtent > len(axes)-i-1 {
			maxExtent = len(axes) - i - 1
		}
		match := true

		for offSet := 0; offSet < maxExtent; offSet++ {
			match = (axes[i-offSet] == axes[i+offSet+1]) && match

			fmt.Println("Col: ", i+1, "MaxExtent: ", maxExtent, "Offset: ", offSet, "Vals: ", axes[i-offSet], axes[i+offSet+1], "Match: ", match)
		}
		if match {
			ans = i + 1
			fmt.Println("Reflection found at col ", i+1)
			break
		}
	}
	return ans
}

func (p *Pattern) fixSmudge() int {
	initReflectionRow := findReflectionLine(p.rows)
	initReflectionCol := findReflectionLine(p.cols)

	for i := 0; i < len(p.rows); i++ {
		for j := 0; j < len(p.cols); j++ {
			p.cols[j] = p.cols[j] ^ (1 << i)
			p.rows[i] = p.rows[i] ^ (1 << j)
			newRow := findReflectionLine(p.rows)
			newCol := findReflectionLine(p.cols)
			if initReflectionRow != newRow {
				return newRow
			}
			if initReflectionCol != newCol {
				return newCol * 100
			}
			p.cols[j] = p.cols[j] ^ (1 << i)
			p.rows[i] = p.rows[i] ^ (1 << j)
		}
	}

	return -1
}

func (p *Pattern) getReflectionValue() int {
	ans := findReflectionLine(p.cols)
	fmt.Println("Reflection found at col ", ans, "Pattern ", p)
	ans += findReflectionLine(p.rows) * 100

	return ans
}

func playPart0(fileName string) int {
	input := readFile(fileName)
	fmt.Println(input)
	patterns := parseForPart0(input)
	fmt.Println(patterns)
	result := 0
	for _, p := range patterns {
		result += p.getReflectionValue()
	}
	return result
}

func playPart1(fileName string) int {
	input := readFile(fileName)
	fmt.Println(input)
	patterns := parseForPart0(input)
	fmt.Println(patterns)
	result := 0
	for _, p := range patterns {
		result += p.fixSmudge()
	}
	return result
}

func parseForPart0(input []string) []*Pattern {
	ans := make([]*Pattern, 0)
	currPattern := &Pattern{}
	ans = append(ans, currPattern)
	for _, line := range input {
		if len(line) == 0 {
			fmt.Println("New Pattern")
			currPattern = &Pattern{}
			ans = append(ans, currPattern)
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
	for idx := 0; idx < len(ans); idx++ {
		pattern := ans[idx]
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

func main() {
	retVal := playPart0("test0.txt")
	fmt.Println(retVal)
	if retVal != 405 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart0("input.txt")
	fmt.Println(retVal)
	if retVal != 43614 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 400 {
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
