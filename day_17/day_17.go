package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Map struct {
	blocks         map[Pos]int
	rowDim, colDim int
}

type Pos [2]int
type Dir [2]int

func (p Pos) moveForward(dir Dir) Pos {
	return Pos{p[0] + dir[0], p[1] + dir[1]}
}

func (p Pos) getDistanceTo(x Pos) int {
	return int(math.Abs(float64(p[0]-x[0])) + math.Abs(float64(p[1]-x[1])))
}

func (p Pos) moveBackward(dir Dir) Pos {
	return Pos{p[0] - dir[0], p[1] - dir[1]}
}

type Path struct {
	pos        Pos
	heatLoss   int
	posHistory []Pos
	dirHistory []Dir
	id         int
}

func (p *Path) isDirValid(dir Dir, maxStraight int) bool {
	if len(p.dirHistory) < maxStraight {
		return true
	}
	ans := false
	for i := 0; i < maxStraight; i++ {
		ans = (p.dirHistory[len(p.dirHistory)-i-1] != dir) || ans
	}
	return ans
}

func (m *Map) printS() string {
	ans := ""
	for row := 0; row < m.rowDim; row++ {
		ans += "\n"
		for col := 0; col < m.colDim; col++ {
			ans += fmt.Sprintf("%d", m.blocks[Pos{row, col}])
		}
	}
	ans += "\n"
	fmt.Println(ans)
	return ans
}

func (m *Map) printPath(p *Path) string {
	ans := ""
	for row := 0; row < m.rowDim; row++ {
		ans += "\n"
		for col := 0; col < m.colDim; col++ {
			currPos := Pos{row, col}
			if p.pos == currPos {
				ans += "X"
				continue
			}
			idx := slices.Index(p.posHistory, currPos)
			if idx > 0 {
				switch p.dirHistory[idx] {
				case Dir{1, 0}:
					ans += "v"
				case Dir{-1, 0}:
					ans += "^"
				case Dir{0, 1}:
					ans += ">"
				case Dir{0, -1}:
					ans += "<"
				}
				continue
			}
			ans += fmt.Sprintf("%d", m.blocks[Pos{row, col}])
		}
	}
	ans += "\n"
	fmt.Println(ans)
	return ans
}

func (m *Map) minimumHeatlLoss(start, end Pos, maxStraight int) int {
	paths := make([]*Path, 0)
	paths = append(paths, &Path{
		pos:        start,
		heatLoss:   0,
		posHistory: make([]Pos, 0),
		dirHistory: make([]Dir, 0),
		id:         0,
	})
	visited := make(map[Pos]map[[3]Pos]int)
	currPos := start
	solutionCandidates := make([]*Path, 0)
	counter := 0
	for len(paths) > 0 {
		counter++
		currPath := paths[0]
		currPos = currPath.pos
		if counter%100000 == 0 {
			fmt.Printf("Counter %d  Num paths %d  Curr Pos %d  Curr Heat Loss %d  Curr Path Length %d Curr len of visitd %d\n", counter, len(paths), currPos, currPath.heatLoss, len(currPath.posHistory), len(visited[currPos]))
		}
		if currPos == end {
			m.printPath(currPath)
			return currPath.heatLoss
		}
		isCurrPosVisited := false
		_, ok := visited[currPos]
		var currPosHistory [3]Pos
		for i := 0; i < len(currPosHistory); i++ {
			if i >= len(currPath.posHistory) {
				currPosHistory[len(currPosHistory)-1-i] = Pos{0, 0}
			} else {
				currPosHistory[len(currPosHistory)-1-i] = currPath.posHistory[len(currPath.posHistory)-1-i]
			}
		}
		if !ok {
			visited[currPos] = make(map[[3]Pos]int)
			visited[currPos][currPosHistory] = currPath.heatLoss
		} else {
			val, ok := visited[currPos][currPosHistory]
			if ok {
				if val <= currPath.heatLoss {
					isCurrPosVisited = true
				}
			} else {
				visited[currPos][currPosHistory] = currPath.heatLoss
			}
		}

		if !isCurrPosVisited {
			for _, dir := range []Dir{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				candPos := currPos.moveForward(dir)
				if candPos[0] < 0 || candPos[0] >= m.rowDim || candPos[1] < 0 || candPos[1] >= m.colDim {
					continue
				}
				if slices.Contains(currPath.posHistory, candPos) {
					continue
				}
				if !currPath.isDirValid(dir, maxStraight) {
					continue
				}
				var currPosHistory [3]Pos
				for i := 1; i < len(currPosHistory); i++ {
					if i >= len(currPath.posHistory) {
						currPosHistory[len(currPosHistory)-1-i] = Pos{0, 0}
					} else {
						currPosHistory[len(currPosHistory)-1-i] = currPath.posHistory[len(currPath.posHistory)-i]
					}
				}
				currPosHistory[2] = currPos
				val, ok := visited[candPos][currPosHistory]
				if ok {
					if val <= currPath.heatLoss+m.blocks[candPos] {
						continue
					}
				}

				newPosHistory := make([]Pos, len(currPath.posHistory))
				copy(newPosHistory, currPath.posHistory)
				newPosHistory = append(newPosHistory, currPos)
				newDirHistory := make([]Dir, len(currPath.dirHistory))
				copy(newDirHistory, currPath.dirHistory)
				newDirHistory = append(newDirHistory, dir)
				paths = append(paths, &Path{
					pos:        candPos,
					heatLoss:   currPath.heatLoss + m.blocks[candPos],
					posHistory: newPosHistory,
					dirHistory: newDirHistory,
					id:         len(paths),
				})
			}
			paths = paths[1:]
			slices.SortFunc(paths, func(a, b *Path) int {
				if a.heatLoss == b.heatLoss {
					return b.id - a.id
				}
				return a.heatLoss - b.heatLoss
			})
		} else {
			paths = paths[1:]
		}
	}
	slices.SortFunc(solutionCandidates, func(a, b *Path) int {
		return b.heatLoss - a.heatLoss
	})

	return solutionCandidates[0].heatLoss
}

func playPart0(fileName string) int {
	input := readFile(fileName)
	fmt.Println(input)
	m := parseForPart0(input)
	m.printS()
	return m.minimumHeatlLoss(Pos{0, 0}, Pos{m.rowDim - 1, m.colDim - 1}, 3)
}

func parseForPart0(input []string) *Map {
	ans := &Map{
		blocks: make(map[Pos]int),
		rowDim: len(input),
		colDim: len(input[0]),
	}
	for row := 0; row < ans.rowDim; row++ {
		for col := 0; col < ans.colDim; col++ {
			ans.blocks[Pos{row, col}], _ = strconv.Atoi(string(input[row][col]))
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
	if retVal != 102 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart0("input.txt")
	fmt.Println(retVal)
	if retVal != 722 {
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
