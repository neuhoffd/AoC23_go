package main

import (
	"fmt"
	"os"
	"strings"
)

func playPart0(fileName string) int {
	input := readFile(fileName)
	fmt.Println(input)
	seq := parseForPart0(input)
	fmt.Println(seq)
	return computeHashSum(seq)
}

func computeHashSum(seq []string) int {
	ans := 0
	for _, s := range seq {
		ans += hash(s)
	}
	return ans
}

func hash(s string) int {
	ans := 0
	for _, c := range s {
		fmt.Println(c, int(c))
		val := int(c)
		ans += val
		ans = ans * 17
		ans = ans % 256
	}
	return ans
}

func parseForPart0(input []string) []string {
	return strings.Split(input[0], ",")
}

func playPart1(fileName string) int {

	return 0
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
