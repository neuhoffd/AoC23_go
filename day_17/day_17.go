package main

import (
	"container/heap"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Path struct {
	pos        Pos
	heatLoss   int
	posHistory []Pos
	dirHistory []Dir
	id         int
}

type State struct {
	pos      Pos
	dir      Dir
	heatLoss int
	straight int
	index    int
	path     *Path
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].heatLoss == pq[j].heatLoss {
		return pq[i].path.id < pq[j].path.id
	}
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	state, _ := x.(*State)
	item := &State{
		path:     state.path,
		pos:      state.pos,
		dir:      state.dir,
		heatLoss: state.heatLoss,
		straight: state.straight,
	}
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	oldQueue := *pq
	n := len(oldQueue)
	item := oldQueue[n-1]
	oldQueue[n-1] = nil
	item.index = -1
	*pq = oldQueue[0 : n-1]
	return item
}

type Map struct {
	blocks         map[Pos]int
	rowDim, colDim int
}

type Pos [2]int
type Dir [2]int

func (p Pos) moveForward(dir Dir) Pos {
	return Pos{p[0] + dir[0], p[1] + dir[1]}
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

func (s State) encode() EncodedState {
	return EncodedState{
		pos:      s.pos,
		dir:      s.dir,
		straight: s.straight,
	}
}

type EncodedState struct {
	pos      Pos
	dir      Dir
	straight int
}

func (m *Map) minimumHeatlLoss(start, end Pos, minStraight, maxStraight int) int {
	counter := 0
	seen := make(map[EncodedState]struct{})

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	initPath := &Path{
		pos:        start,
		heatLoss:   0,
		posHistory: make([]Pos, 0),
		dirHistory: make([]Dir, 0),
		id:         0,
	}
	initState := &State{
		pos:      start,
		dir:      [2]int{},
		heatLoss: 0,
		straight: 0,
		index:    0,
		path:     initPath,
	}

	heap.Push(&pq, initState)

	for len(pq) > 0 {
		currState := heap.Pop(&pq).(*State)
		currPos := currState.pos
		currPath := currState.path
		if currPos == end && currState.straight >= minStraight-1 {
			m.printPath(currPath)
			return currPath.heatLoss
		}
		if _, isSeen := seen[currState.encode()]; isSeen {
			continue
		}
		seen[currState.encode()] = struct{}{}
		for _, dir := range []Dir{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			counter++
			newStraight := 0
			if dir == currState.dir {
				newStraight = currState.straight + 1
			} else {
				if currState.straight < minStraight-1 && (dir != currState.dir && currState.dir != Dir{0, 0}) {
					continue
				}
			}
			if newStraight > maxStraight-1 {
				continue
			}
			candPos := currPos.moveForward(dir)
			if candPos[0] < 0 || candPos[0] >= m.rowDim || candPos[1] < 0 || candPos[1] >= m.colDim || slices.Contains(currPath.posHistory, candPos) {
				continue
			}
			newPosHistory := make([]Pos, len(currPath.posHistory))
			copy(newPosHistory, currPath.posHistory)
			newPosHistory = append(newPosHistory, currPos)
			newDirHistory := make([]Dir, len(currPath.dirHistory))
			copy(newDirHistory, currPath.dirHistory)
			newDirHistory = append(newDirHistory, dir)
			newState := State{
				pos:      candPos,
				dir:      dir,
				heatLoss: currPath.heatLoss + m.blocks[candPos],
				straight: newStraight,
				path: &Path{
					pos:        candPos,
					heatLoss:   currPath.heatLoss + m.blocks[candPos],
					posHistory: newPosHistory,
					dirHistory: newDirHistory,
					id:         counter,
				}}
			if _, isSeen := seen[newState.encode()]; isSeen {
				continue
			}
			heap.Push(&pq, &newState)
		}
	}
	return -1
}

func playPart0(fileName string) int {
	input := readFile(fileName)
	fmt.Println(input)
	m := parseForPart0(input)
	m.printS()
	return m.minimumHeatlLoss(Pos{0, 0}, Pos{m.rowDim - 1, m.colDim - 1}, 0, 3)
}

func playPart1(fileName string) int {
	input := readFile(fileName)
	fmt.Println(input)
	m := parseForPart0(input)
	m.printS()
	return m.minimumHeatlLoss(Pos{0, 0}, Pos{m.rowDim - 1, m.colDim - 1}, 4, 10)
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
	if retVal != 94 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != 894 {
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
