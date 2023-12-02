package main

import (
	"fmt"
	"os"
	"strings"
)

func parseForPart1(fileName string) [][]int {

	return nil
}

func parseForPart2(fileName string) [][]int {
	return nil
}

func readFile(fileName string) []string {
	bytes, _ := os.ReadFile(fileName)
	input := string(bytes)
	splitInput := strings.Split(input, "\n")
	return splitInput
}

func playPart1(fileName string) int {
	lines := readFile(fileName)

	return 0
}

func playPart2(fileName string) int {

	return 0
}

func main() {
	retVal := playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 8 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != 54697 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart2("test1.txt")
	fmt.Println(retVal)
	if retVal != 281 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = playPart2("input.txt")
	fmt.Println(retVal)
	if retVal != 54885 {
		panic("Part 1 failed")
	}
	fmt.Println("Part 1 passed")
}
