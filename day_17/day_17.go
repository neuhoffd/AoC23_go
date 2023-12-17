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
	})
	visited := make(map[Pos]map[int]int)
	currPos := start
	solutionCandidates := make([]*Path, 0)
	counter := 0
	for len(paths) > 0 {
		counter++
		currPath := paths[0]
		currPos = currPath.pos
		/*alarmPos := Pos{1, 5}
		if currPos == alarmPos {
			fmt.Println("obacht")
		}*/
		if counter%1000 == 0 {
			fmt.Printf("Counter %d  Num paths %d  Curr Pos %d  Curr Heat Loss %d  Curr Path Length %d\n", counter, len(paths), currPos, currPath.heatLoss, len(currPath.posHistory))
			//m.printPath(currPath)
		}
		//m.printPath(currPath)
		if len(currPath.posHistory) > 0 && currPath.posHistory[len(currPath.posHistory)-1].getDistanceTo(currPos) > 1 {
			fmt.Println("This shouldn't happen")
		}
		//m.printPath(currPath)
		if currPos == end {
			m.printPath(currPath)
			return currPath.heatLoss
		}
		shortestPath, ok := visited[currPos][currPath.heatLoss]
		isCurrPosVisited := false

		isCurrPosVisited = ok && len(currPath.posHistory) >= shortestPath
		if isCurrPosVisited {
			//fmt.Println()
		}

		if !ok || !isCurrPosVisited {
			for _, dir := range []Dir{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				candPos := currPos.moveForward(dir)
				if candPos.getDistanceTo(currPos) > 1 {
					fmt.Println("This shouldn't happen")
				}

				if candPos[0] < 0 || candPos[0] >= m.rowDim || candPos[1] < 0 || candPos[1] >= m.colDim {
					continue
				}
				if slices.Contains(currPath.posHistory, candPos) {
					continue
				}
				if !currPath.isDirValid(dir, maxStraight) {
					continue
				}
				if candPos.getDistanceTo(currPos) > 1 {
					fmt.Println("This shouldn't happen")
				}
				if len(currPath.posHistory) > 0 && currPath.posHistory[len(currPath.posHistory)-1].getDistanceTo(candPos) > 2 {
					fmt.Println("This shouldn't happen")
				}
				shortestPath, ok := visited[candPos][currPath.heatLoss+m.blocks[candPos]]
				if ok {
					if shortestPath < len(currPath.posHistory)+1 {
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
				})

			}
			shortestPath, ok = visited[currPos][currPath.heatLoss]
			if ok {
				if shortestPath > len(currPath.posHistory) {
					visited[currPos][currPath.heatLoss] = len(currPath.posHistory)
				}
			} else {
				if len(visited[currPos]) == 0 {
					visited[currPos] = make(map[int]int)
				}
				visited[currPos][currPath.heatLoss] = len(currPath.posHistory)
			}
			paths = paths[1:]
			slices.SortFunc(paths, func(a, b *Path) int {
				if a.heatLoss == b.heatLoss {
					return len(a.posHistory) - len(b.posHistory)
				}
				return a.heatLoss - b.heatLoss
			})
		} else {
			paths = paths[1:]
			//fmt.Printf("Letting path %+v die\n", currPath)
		}
	}
	slices.SortFunc(solutionCandidates, func(a, b *Path) int {
		return b.heatLoss - a.heatLoss
	})

	return solutionCandidates[0].heatLoss
}

func (m *Map) minimumHeatlLossOld(start, end Pos, maxStraight int) int {
	distances := make(map[Pos]int)
	paths := make([]*Path, 0)
	fmt.Println(paths)
	pathDirections := make(map[Pos]Dir)
	pathDirections[start] = Dir{1, 1}
	unvisited := make([]Pos, 0)
	for k := range m.blocks {
		distances[k] = -1
		if k != start {
			unvisited = append(unvisited, k)
		}
	}
	distances[start] = 0
	currPos := start
	visited := make([]Pos, 0)

	for len(visited) <= len(m.blocks) {
		visited = append(visited, currPos)
		fmt.Printf("Curr Pos %d\n", currPos)
		//m.printPath(pathDirections)
		if currPos == end {
			return distances[currPos]
		}
		neighbors := make([]Pos, 0)
		for _, dir := range []Dir{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			candPos := currPos.moveForward(dir)
			if candPos[0] < 0 || candPos[0] >= m.rowDim || candPos[1] < 0 || candPos[1] >= m.colDim || slices.Contains(visited, candPos) {
				continue
			}
			if dir == pathDirections[currPos] {
				cursor := currPos
				currDir := dir
				valid := false
				for i := 0; i < maxStraight-1; i++ {
					cursor = Pos{cursor[0] - pathDirections[cursor][0], cursor[1] - pathDirections[cursor][1]}
					fmt.Println(cursor, pathDirections[cursor])
					if pathDirections[cursor] != currDir {
						valid = true
						break
					}
					currDir = pathDirections[cursor]
				}
				if !valid {
					continue
				}
			} else {
				if dir[0] == (-1)*pathDirections[currPos][0] || dir[1] == (-1)*pathDirections[currPos][1] {
					continue
				}
			}

			if pathDirections[currPos] == pathDirections[candPos] &&
				pathDirections[Pos{currPos[0] - pathDirections[currPos][0], currPos[1] - pathDirections[currPos][1]}] == pathDirections[candPos] {
				fmt.Println("Curr ", currPos, "Cand ", candPos)
			}
			if distances[candPos] < 0 || distances[candPos] > distances[currPos]+m.blocks[candPos] {
				distances[candPos] = distances[currPos] + m.blocks[candPos]

				pathDirections[candPos] = dir
			}
			neighbors = append(neighbors, candPos)
		}
		if len(neighbors) > 0 {
			currPos = slices.MinFunc(neighbors, func(a, b Pos) int {
				return distances[a] - distances[b]
			})
		} else {
			currPos = slices.MinFunc(unvisited, func(a, b Pos) int {
				if distances[a] == -1 {
					return 1
				}
				if distances[b] == -1 {
					return -1
				}
				return distances[a] - distances[b]
			})
		}
	}

	return distances[end]
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
