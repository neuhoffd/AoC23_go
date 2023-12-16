package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Operation struct {
	label string
	value int
	op    string
	s     string
}

type Box struct {
	contents []Operation
}

func computeHashSum(seq []Operation) int {
	ans := 0
	for _, s := range seq {
		ans += hash(s.s)
	}
	return ans
}

func hash(s string) int {
	ans := 0
	for _, c := range s {
		val := int(c)
		ans = ((ans + val) * 17) % 256
	}
	return ans
}

func runInitSequence(seq []Operation) []*Box {
	boxes := make([]*Box, 256)
	for i := 0; i < len(boxes); i++ {
		boxes[i] = &Box{contents: make([]Operation, 0)}
	}
	for idx, op := range seq {
		fmt.Println("Op #", idx, ", ", op)
		boxId := hash(op.label)
		lensIdx := slices.IndexFunc(boxes[boxId].contents, func(lens Operation) bool { return lens.label == op.label })
		if op.op == "-" {
			if lensIdx >= 0 {
				boxes[boxId].contents = append(boxes[boxId].contents[:lensIdx], boxes[boxId].contents[lensIdx+1:]...)
			}
		} else {
			if lensIdx >= 0 {
				boxes[boxId].contents[lensIdx].value = op.value
			} else {
				boxes[boxId].contents = append(boxes[boxId].contents, op)
			}
		}
	}
	return boxes
}

func computeFocusingPowerSum(boxes []*Box) int {
	ans := 0
	for i := 0; i < len(boxes); i++ {
		for j := 0; j < len(boxes[i].contents); j++ {
			ans += (i + 1) * (j + 1) * boxes[i].contents[j].value
		}
	}
	return ans
}

func parseForPart0(input []string) []Operation {
	r := regexp.MustCompile(`-|=`)
	splits := strings.Split(input[0], ",")
	ans := make([]Operation, len(splits))
	for idx, s := range splits {
		ans[idx].s = s
		cmdStrings := r.Split(s, -1)
		ans[idx].label = cmdStrings[0]
		ans[idx].value, _ = strconv.Atoi(cmdStrings[1])
		ans[idx].op = r.FindAllString(s, -1)[0]
	}
	return ans
}

func playPart0(fileName string) int {
	input := readFile(fileName)
	seq := parseForPart0(input)
	return computeHashSum(seq)
}

func playPart1(fileName string) int {
	input := readFile(fileName)
	seq := parseForPart0(input)
	boxes := runInitSequence(seq)
	return computeFocusingPowerSum(boxes)
}

func main() {
	retVal := playPart0("test0.txt")
	fmt.Println(retVal)
	if retVal != 1320 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart0("input.txt")
	fmt.Println(retVal)
	if retVal != 520500 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 145 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != 213097 {
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
