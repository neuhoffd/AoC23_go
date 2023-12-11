package main

import (
	"fmt"
	"os"
	"strings"
)

type Galaxy struct {
	row, col, id int
}

type Universe struct {
	galaxies             []*Galaxy
	emptyRows, emptyCols []int
	rows, cols           int
}

func playPart0(fileName string) int {
	lines := readFile(fileName)
	fmt.Println(lines)
	universe := parseForPart0(lines)
	fmt.Println(universe)
	return 0
}

func parseForPart0(input []string) *Universe {
	rows := len(input)
	cols := len(input[0])
	ans := &Universe{
		rows: rows,
		cols: cols,
	}
	galaxiesFoundbyCol := make([]int, cols)
	for row := 0; row < rows; row++ {
		galaxiesFoundRow := 0
		for col := 0; col < cols; col++ {
			if input[row][col] == '#' {
				newGalaxy := &Galaxy{
					row: row,
					col: col,
					id:  len(ans.galaxies) + 1,
				}
				ans.galaxies = append(ans.galaxies, newGalaxy)
				galaxiesFoundRow++
				galaxiesFoundbyCol[col]++
			}
		}
		if galaxiesFoundRow == 0 {
			ans.emptyRows = append(ans.emptyRows, row)
		}
	}
	for col := 0; col < cols; col++ {
		if galaxiesFoundbyCol[col] == 0 {
			ans.emptyCols = append(ans.emptyCols, col)
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
	if retVal != 374 {
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
