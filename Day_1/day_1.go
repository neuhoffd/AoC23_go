package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseFile(fileName string) [][]int {
	input := readFile(fileName)
	fmt.Println(input)

	var numbers [][]int
	for _, line := range input {
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
	numberStrings := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

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

func playPart1(fileName string) int {
	input := parseFile(fileName)
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

	/*lines = parseFile("input.txt", false)
	// fmt.Println(lines)
	score, _ = resolve_lines(lines, 2)
	fmt.Println(score)
	if score != 5576 {
		panic("Part 1 failed")
	}

	lines = parseFile("test.txt", true)
	fmt.Println(lines)
	score, coords = resolve_lines(lines, 2)
	fmt.Println(coords)
	fmt.Println(score)
	if score != 12 {
		panic("Test 2 failed")
	}

	lines = parseFile("input.txt", true)
	fmt.Println(lines)
	score, _ = resolve_lines(lines, 2)
	fmt.Println(score)
	if ssumcore != 18144 {
		panic("Part 2 failed")
	}*/
}
