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

func (r *Record) dp(i, j int, cache [][]int) int {
	if i >= len(r.record) {
		if j < len(r.groups) {
			return 0
		}
		return 1
	}
	if cache[i][j] != -1 {
		return cache[i][j]
	}
	res := 0
	if r.record[i] == '.' {
		res += r.dp(i+1, j, cache)
	} else {
		if r.record[i] == '?' {
			res += r.dp(i+1, j, cache)
		}
		if j < len(r.groups) {
			damagedCount := byte(0)
			for k := i; k < len(r.record); k++ {
				if r.record[k] == '.' || (r.groups[j] == damagedCount && r.record[k] == '?') || damagedCount > r.groups[j] {
					break
				}
				damagedCount++
			}
			if damagedCount == r.groups[j] {
				if i+int(damagedCount) < len(r.record) && r.record[i+int(damagedCount)] != '#' {
					res += r.dp(i+int(damagedCount)+1, j+1, cache)
				} else {
					res += r.dp(i+int(damagedCount), j+1, cache)
				}
			}
		}
	}
	cache[i][j] = res
	return res
}

func (r *Record) countArrangements() int {
	var cache [][]int
	for i := 0; i < len(r.record); i++ {
		cache = append(cache, make([]int, len(r.groups)+1))
		for j := 0; j < len(r.groups)+1; j++ {
			cache[i][j] = -1
		}
	}

	return r.dp(0, 0, cache)
}

func playPart0(fileName string) int {
	input := readFile(fileName)
	fmt.Println(input)
	records := parseForPart0(input)
	fmt.Printf("Records\n%+v\n", records)
	result := 0
	for _, r := range records {
		newVal := r.countArrangements()
		result += newVal
		fmt.Printf("Record: %+v\nCount: %d\nResult: %d\n", *r, newVal, result)
	}
	return result
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
