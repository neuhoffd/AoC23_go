package main

import (
	"fmt"
	"os"
	"strings"
)

type Pattern struct {
	rows, cols []int
	nCols      int
	line       int
}

func findReflectionLine(axes []int, ignore int) int {
	ans := -1
	for i := 0; i < len(axes)-1; i++ {
		maxExtent := i + 1
		if maxExtent > len(axes)-i-1 {
			maxExtent = len(axes) - i - 1
		}
		match := true
		for offSet := 0; offSet < maxExtent; offSet++ {
			match = (axes[i-offSet] == axes[i+offSet+1]) && match
		}
		if match && i+1 != ignore {
			ans = i + 1
			return ans
		}
	}
	return ans
}

func (p *Pattern) fixSmudge() int {
	initReflectionRow := findReflectionLine(p.rows, -1)
	initReflectionCol := findReflectionLine(p.cols, -1)
	ans := 0
	for i := 0; i < len(p.rows); i++ {
		for j := 0; j < len(p.cols); j++ {
			p.cols[j] = p.cols[j] ^ (1 << (len(p.rows) - i - 1))
			p.rows[i] = p.rows[i] ^ (1 << (len(p.cols) - j - 1))
			newRow := findReflectionLine(p.rows, initReflectionRow)
			newCol := findReflectionLine(p.cols, initReflectionCol)
			if initReflectionRow != newRow && newRow > 0 {

				return newRow * 100
			}
			if initReflectionCol != newCol && newCol > 0 {
				return newCol
			}
			p.cols[j] = p.cols[j] ^ (1 << (len(p.rows) - i - 1))
			p.rows[i] = p.rows[i] ^ (1 << (len(p.cols) - j - 1))
		}
	}

	return ans
}

func (p *Pattern) getReflectionValue() int {
	ans := findReflectionLine(p.cols, -1)
	if ans > 0 {
		return ans
	}
	ans = findReflectionLine(p.rows, -1) * 100
	if ans > 0 {
		return ans
	}
	if ans <= 0 {
		return 0
	}
	return -1
}

func playPart0(fileName string) int {
	input := readFile(fileName)
	patterns := parseForPart0(input)
	result := 0
	for _, p := range patterns {
		result += p.getReflectionValue()
	}
	return result
}

func playPart1(fileName string) int {
	input := readFile(fileName)
	patterns := parseForPart0(input)
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
	for idx, line := range input {
		if len(line) == 0 {
			currPattern = &Pattern{}
			currPattern.line = idx + 1
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
			}
		}
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
	if retVal != 36771 {
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
