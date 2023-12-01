package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseForPart1(fileName string) [][]int {
	input := readFile(fileName)
	fmt.Println(input)

	var numbers [][]int
	for _, line := range input {
		numbers = append(numbers, extractNumbers(line))
	}

	return numbers
}

func parseForPart2(fileName string) [][]int {
	input := readFile(fileName)
	fmt.Println(input)

	var numbers [][]int
	for _, line := range input {
		line = replaceNumberStrings(line)
		numbers = append(numbers, extractNumbers(line))
	}

	return numbers
}

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR!", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var ans []string
	for scanner.Scan() {
		line := scanner.Text()
		ans = append(ans, line)
	}
	return ans
}

func extractNumbers(line string) []int {
	fmt.Println("Extracting numbers from ", line)
	var ans []int
	for i := 0; i < len(line); i++ {
		fmt.Printf("%c \n", line[i])
		val, err := strconv.Atoi(string(line[i]))
		if err == nil {
			ans = append(ans, val)
		}
	}
	return ans
}

func replaceNumberStrings(line string) string {
	numberStrings := map[string]string{
		"zero":  "z0o",
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	fmt.Println("Replacing number strings in ", line)
	for k, v := range numberStrings {
		line = strings.ReplaceAll(line, k, v)
	}
	return line
}

func playPart1(fileName string) int {
	input := parseForPart1(fileName)
	fmt.Println(input)

	ans := 0
	for _, line := range input {
		if len(line) == 0 {
			panic("No numbers found, exciting")
		}
		ans += line[0]*10 + line[len(line)-1]
	}

	return ans
}

func playPart2(fileName string) int {
	input := parseForPart2(fileName)
	fmt.Println(input)

	ans := 0
	for _, line := range input {
		if len(line) == 0 {
			panic("No numbers found, exciting")
		}
		ans += line[0]*10 + line[len(line)-1]
	}

	return ans
}

func main() {
	retVal := playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 142 {
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
