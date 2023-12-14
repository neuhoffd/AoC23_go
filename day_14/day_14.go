package main

import (
	"fmt"
	"os"
	"strings"
)

type Platform struct {
	p [][]string
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

func (p *Platform) turnRight() {
	np := Platform{}
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

func playPart0(fileName string) int {
	input := readFile(fileName)
	fmt.Println(input)
	platform := parseForPart0(input)
	platform.print()
	platform.rollNorth()
	platform.print()
	return platform.getLoad()
}

func parseForPart0(input []string) *Platform {
	ans := &Platform{}
	for _, line := range input {
		*ans = append(*ans, strings.Split(line, ""))
	}
	return ans
}

func playPart1(fileName string) int {

	return 0
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
