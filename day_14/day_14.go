package main

import (
	"fmt"
	"os"
	"strings"
)

type Platform struct {
	p    [][]string
	seen map[string]int
}

func (p *Platform) print() {
	fmt.Printf("\n")
	for _, row := range p.p {
		for _, col := range row {
			fmt.Printf("%s", col)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (p *Platform) rollNorth() {
	for j := 0; j < len(p.p[0]); j++ {
		for i := 0; i < len(p.p); i++ {
			if p.p[i][j] == "O" {
				idx := i
				for idx > 0 {
					if p.p[idx-1][j] == "." {
						p.p[idx-1][j] = "O"
						p.p[idx][j] = "."
						idx--
					} else {
						break
					}
				}
			}
		}
	}
}

func (p *Platform) toVector() []string {
	ans := make([]string, 0)
	for _, str := range p.p {
		ans = append(ans, str...)
	}
	return ans
}

func (p *Platform) turnRight() {
	np := make([][]string, len(p.p[0]))
	for i := 0; i < len(np); i++ {
		np[i] = make([]string, len(p.p))
	}
	for i := 0; i < len(p.p); i++ {
		for j := 0; j < len(p.p[0]); j++ {
			np[j][len(p.p)-1-i] = p.p[i][j]
		}
	}
	p.p = np
}

func (p *Platform) getLoad() int {
	ans := 0
	for i := 0; i < len(p.p); i++ {
		currLoadValue := len(p.p) - i
		for j := 0; j < len(p.p[0]); j++ {
			if p.p[i][j] == "O" {
				ans += currLoadValue
			}
		}
	}
	return ans
}

func parseForPart0(input []string) *Platform {
	ans := &Platform{seen: make(map[string]int)}
	for _, line := range input {
		ans.p = append(ans.p, strings.Split(line, ""))
	}
	return ans
}

func playPart0(fileName string) int {
	input := readFile(fileName)
	platform := parseForPart0(input)
	platform.print()
	platform.rollNorth()
	return platform.getLoad()
}

func playPart1(fileName string) int {
	input := readFile(fileName)
	platform := parseForPart0(input)
	platform.print()
	cycles := 0
	cycleFound := false
	for cycles < 1000000000 {
		key := strings.Join(platform.toVector(), "")
		val, ok := platform.seen[key]
		phase := -1
		if ok && !cycleFound {
			phase = cycles - val
		} else {
			phase = 0
		}
		if phase > 0 {
			cycles = 1000000000 - (1000000000-cycles)%phase
			cycleFound = true
			continue
		}
		platform.seen[key] = cycles
		if cycles%1000000 == 0 {
			fmt.Printf("Cycles: %d, %f of 1000000000\n", cycles, float64(cycles/10000000))
		}
		for i := 0; i < 4; i++ {
			platform.rollNorth()
			platform.turnRight()
		}
		cycles++
	}
	return platform.getLoad()
}

func main() {
	retVal := playPart0("test0.txt")
	fmt.Println(retVal)
	if retVal != 136 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart0("input.txt")
	fmt.Println(retVal)
	if retVal != 108889 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 64 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != 104671 {
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
