package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Card struct {
	number       int
	winning, own []int
}

func parseForPart1(input []string) []Card {
	cards := []Card{}

	for _, line := range input {
		r := regexp.MustCompile(`\s|\|`)
		splits := r.Split(line, -1)
		fmt.Println(splits)
	}
	return []Card{}
}

func readFile(fileName string) []string {
	bytes, _ := os.ReadFile(fileName)
	input := string(bytes)
	splitInput := strings.Split(input, "\n")
	return splitInput
}

func playPart1(fileName string) int {
	input := readFile(fileName)

	fmt.Println(input)
	return 0
}

func playPart2(fileName string) int {

	return 0
}

func main() {
	retVal := playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 13 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != 0 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart2("test0.txt")
	fmt.Println(retVal)
	if retVal != 0 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = playPart2("input.txt")
	fmt.Println(retVal)
	if retVal != 0 {
		panic("Part 1 failed")
	}
	fmt.Println("Part 1 passed")
}
