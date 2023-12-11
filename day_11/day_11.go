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
	expandBy             int
}

func (uv *Universe) computeExpansionAddition(g1, g2 *Galaxy) int {
	minRow := slices.Min([]int{g1.row, g2.row})
	maxRow := slices.Max([]int{g1.row, g2.row})
	minCol := slices.Min([]int{g1.col, g2.col})
	maxCol := slices.Max([]int{g1.col, g2.col})

	ans := 0

	for row := minRow + 1; row < maxRow; row++ {
		if uv.emptyRows[row] {
			ans += uv.expandBy
		}
	}

	for col := minCol + 1; col < maxCol; col++ {
		if uv.emptyCols[col] {
			ans += uv.expandBy
		}
	}
	return ans
}

func (uv *Universe) getSumOfShortestPathsBwGalaxies() int {
	slices.SortFunc(uv.galaxies, func(a, b *Galaxy) int { return a.id - b.id })
	ans := 0

	for a := 0; a < len(uv.galaxies); a++ {
		for b := a + 1; b < len(uv.galaxies); b++ {
			currDist := 0
			currDist += uv.galaxies[a].computeManhattanDistance(uv.galaxies[b])
			currDist += uv.computeExpansionAddition(uv.galaxies[a], uv.galaxies[b])
			ans += currDist
		}
	}
	return ans
}

func (g1 *Galaxy) computeManhattanDistance(g2 *Galaxy) int {
	return int(math.Abs(float64(g1.row)-float64(g2.row)) + math.Abs(float64(g1.col)-float64(g2.col)))
}

func play(fileName string, expandBy int) int {
	lines := readFile(fileName)
	universe := parseForPart0(lines)
	universe.expandBy = expandBy
	return universe.getSumOfShortestPathsBwGalaxies()
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

func main() {
	retVal := play("test0.txt", 1)
	fmt.Println(retVal)
	if retVal != 374 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = play("input.txt", 1)
	fmt.Println(retVal)
	if retVal != 10033566 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = play("test0.txt", 99)
	fmt.Println(retVal)
	if retVal != 8410 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = play("input.txt", 999999)
	fmt.Println(retVal)
	if retVal != 560822911938 {
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
