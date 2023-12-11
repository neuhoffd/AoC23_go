package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Galaxy struct {
	row, col, id int
}

type Universe struct {
	galaxies             []*Galaxy
	emptyRows, emptyCols []bool
	rows, cols           int
}

func (uv *Universe) computeExpansionAddition(g1, g2 *Galaxy) int {
	//need to compute

	return 0
}

func (uv *Universe) getSumOfShortestPathsBwGalaxies() int {
	slices.SortFunc(uv.galaxies, func(a, b *Galaxy) int { return a.id - b.id })
	ans := 0

	for a := 0; a < len(uv.galaxies); a++ {
		for b := a; b < len(uv.galaxies); b++ {
			ans += computeManhattanDistance(uv.galaxies[a], uv.galaxies[b])
		}
	}
	return ans
}

func computeManhattanDistance(g1, g2 *Galaxy) int {
	return int(math.Abs(float64(g1.row) - float64(g2.row) + math.Abs(float64(g1.col)-float64(g2.col))))
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
		rows:      rows,
		cols:      cols,
		emptyRows: make([]bool, rows),
		emptyCols: make([]bool, cols),
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
		ans.emptyRows[row] = galaxiesFoundRow == 0
	}
	for col := 0; col < cols; col++ {
		ans.emptyCols[col] = galaxiesFoundbyCol[col] == 0
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
