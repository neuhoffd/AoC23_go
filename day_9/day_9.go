package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type History struct {
	sequences [][]int
}

func (hist *History) determineSequences() {
	i := 0
	for !(slices.Max(hist.sequences[i]) == 0 && slices.Min(hist.sequences[i]) == 0) {
		var newSequence []int
		for idx := 0; idx < len(hist.sequences[i])-1; idx++ {
			newSequence = append(newSequence, hist.sequences[i][idx+1]-hist.sequences[i][idx])
		}
		hist.sequences = append(hist.sequences, newSequence)
		i++
	}
}

func (hist *History) extrapolatePart1() int {
	hist.determineSequences()
	hist.sequences[len(hist.sequences)-1] = append(hist.sequences[len(hist.sequences)-1], 0)
	for i := len(hist.sequences) - 2; i >= 0; i-- {
		hist.sequences[i] = append(hist.sequences[i],
			hist.sequences[i][len(hist.sequences[i])-1]+hist.sequences[i+1][len(hist.sequences[i+1])-1])
	}
	return hist.sequences[0][len(hist.sequences[0])-1]
}

func (hist *History) extrapolatePart2() int {
	hist.determineSequences()
	hist.sequences[len(hist.sequences)-1] = append(hist.sequences[len(hist.sequences)-1], 0)
	for i := len(hist.sequences) - 2; i >= 0; i-- {
		hist.sequences[i] = append([]int{
			hist.sequences[i][0] - hist.sequences[i+1][0]},
			hist.sequences[i]...)
	}
	return hist.sequences[0][0]
}

func playPart1(fileName string) int {
	input := readFile(fileName)
	histories := parseForPart1(input)
	result := int(0)
	for i := 0; i < len(histories); i++ {
		result += histories[i].extrapolatePart1()
	}
	return result
}

func playPart2(fileName string) int {
	input := readFile(fileName)
	histories := parseForPart1(input)
	result := int(0)
	for i := 0; i < len(histories); i++ {
		result += histories[i].extrapolatePart2()
	}
	return result
}

func parseForPart1(input []string) []History {
	r := regexp.MustCompile(`\-?\d+`)
	ans := make([]History, 0)
	for _, line := range input {
		splits := r.FindAllString(line, -1)
		currSeq := make([]int, len(splits))
		for i := 0; i < len(currSeq); i++ {
			currVal, _ := strconv.Atoi(splits[i])
			currSeq[i] = int(currVal)
		}
		ans = append(ans, History{
			sequences: [][]int{currSeq},
		})
	}
	return ans
}

func main() {
	retVal := playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 114 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != 1877825184 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart2("test0.txt")
	fmt.Println(retVal)
	if retVal != 2 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = playPart2("input.txt")
	fmt.Println(retVal)
	if retVal != 1108 {
		panic("Part 1 failed")
	}
	fmt.Println("Part 1 passed")
}

func readFile(fileName string) []string {
	bytes, _ := os.ReadFile(fileName)
	input := string(bytes)
	splitInput := strings.Split(input, "\n")
	return splitInput
}
