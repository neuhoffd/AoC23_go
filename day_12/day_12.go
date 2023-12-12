package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	record []byte
	groups []byte
}

func playPart0(fileName string) int {
	input := readFile(fileName)
	fmt.Println(input)
	records := parseForPart0(input)
	fmt.Printf("Records\n%+v\n", records)
	return 0
}

func parseForPart0(input []string) []*Record {
	ans := make([]*Record, 0)

	for _, line := range input {
		newRecord := &Record{}
		splits := strings.Fields(line)
		for _, char := range strings.Split(splits[0], "") {
			newRecord.record = append(newRecord.record, char[0])
		}
		for _, num := range strings.Split(splits[1], ",") {
			val, _ := strconv.Atoi(num)
			newRecord.groups = append(newRecord.groups, byte(val))
		}
		ans = append(ans, newRecord)
	}
	return ans
}

func playPart1(fileName string) int {

	return 0
}

func main() {
	retVal := playPart0("test0.txt")
	fmt.Println(retVal)
	if retVal != 21 {
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
