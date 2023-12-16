package main

import (
	"fmt"
	"os"
	"strings"
)

type Beam struct {
	dir, pos []int
}

type Tile struct {
	pos       []int
	symbol    string
	energized bool
	visited   [][]int
}

func printS(g [][]Tile) string {
	ans := ""
	for _, row := range g {
		for _, tile := range row {
			str := tile.symbol
			if tile.energized {
				str = "#"
			}
			ans += fmt.Sprintf("%s ", str)
		}
		ans += fmt.Sprintf("\n")
	}
	fmt.Println(ans)
	return ans
}

func countEnergized(g [][]Tile) int {
	ans := 0
	for _, row := range g {
		for _, tile := range row {
			if tile.energized {
				ans++
			}
		}
	}
	return ans
}

func shootBeam(g [][]Tile) int {
	ans := 0
	beams := []Beam{
		Beam{dir: []int{0, 1},
			pos: []int{0, 0}},
	}
	steps := 0
	for len(beams) > 0 {
		for i := 0; i < len(beams); i++ {
			//fmt.Println(beams)
			if beams[i].pos[0] >= len(g) || beams[i].pos[1] >= len(g[0]) || beams[i].pos[0] < 0 || beams[i].pos[1] < 0 {
				beams = append(beams[:i], beams[i+1:]...)
				continue
			}

			if !g[beams[i].pos[0]][beams[i].pos[1]].energized {
				ans++
				g[beams[i].pos[0]][beams[i].pos[1]].energized = true
			}
			/*if slices.ContainsFunc(g[beams[i].pos[0]][beams[i].pos[1]].visited, func(inp []int) bool { return inp[0] == beams[i].dir[0] && inp[1] == beams[i].dir[1] }) {
				beams = append(beams[:i], beams[i+1:]...)
				continue
			}*/
			if len(g[beams[i].pos[0]][beams[i].pos[1]].visited) > 1000 {
				beams = append(beams[:i], beams[i+1:]...)
				continue
			}
			g[beams[i].pos[0]][beams[i].pos[1]].visited = append(g[beams[i].pos[0]][beams[i].pos[1]].visited, beams[i].dir)
			switch g[beams[i].pos[0]][beams[i].pos[1]].symbol {
			case "\\":
				{
					if beams[i].dir[0] != 0 {
						beams[i].dir[1] = beams[i].dir[0]
						beams[i].dir[0] = 0
					} else {
						beams[i].dir[0] = beams[i].dir[1]
						beams[i].dir[1] = 0
					}
				}
			case "/":
				{
					if beams[i].dir[0] != 0 {
						beams[i].dir[1] = (-1) * beams[i].dir[0]
						beams[i].dir[0] = 0
					} else {
						beams[i].dir[0] = (-1) * beams[i].dir[1]
						beams[i].dir[1] = 0
					}
				}
			case "|":
				{
					if beams[i].dir[1] != 0 {
						beams = append(beams, Beam{
							pos: []int{beams[i].pos[0], beams[i].pos[1]},
							dir: []int{(-1) * beams[i].dir[1], 0},
						})
						beams[len(beams)-1].pos[0] += beams[len(beams)-1].dir[0]
						beams[len(beams)-1].pos[1] += beams[len(beams)-1].dir[1]

						beams[i].dir[0] = beams[i].dir[1]
						beams[i].dir[1] = 0
					}
				}
			case "-":
				{
					if beams[i].dir[0] != 0 {
						beams = append(beams, Beam{
							pos: []int{beams[i].pos[0], beams[i].pos[1]},
							dir: []int{0, (-1) * beams[i].dir[0]},
						})
						beams[len(beams)-1].pos[0] += beams[len(beams)-1].dir[0]
						beams[len(beams)-1].pos[1] += beams[len(beams)-1].dir[1]

						beams[i].dir[1] = beams[i].dir[0]
						beams[i].dir[0] = 0
					}
				}
			}
			beams[i].pos[0] += beams[i].dir[0]
			beams[i].pos[1] += beams[i].dir[1]
			//_ = printS(g)
			//fmt.Println(countEnergized(g))
		}
		steps++
	}
	//fmt.Println(ans, countEnergized(g))

	return ans
}

func playPart0(fileName string) int {
	input := readFile(fileName)
	tiles := parseForPart0(input)
	return shootBeam(tiles)
}

func parseForPart0(input []string) [][]Tile {
	ans := make([][]Tile, len(input))
	for row := 0; row < len(input); row++ {
		ans[row] = make([]Tile, len(input[row]))
		for col := 0; col < len(input[row]); col++ {
			ans[row][col] = Tile{
				pos:       []int{row, col},
				symbol:    string(input[row][col]),
				energized: false,
			}
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
	if retVal != 46 {
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
