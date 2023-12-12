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

func (r *Record) dp(i, j int, c [][]int) int {
	if i >= len(r.record) {
		if j < len(r.groups) {
			return 0
		}
		return 1
	}
	if c[i][j] != -1 {
		return c[i][j]
	}
	res := 0
	if r.record[i] == '.' {
		res += r.dp(i+1, j, c)
	} else {
		if r.record[i] == '?' {
			res += r.dp(i+1, j, c)
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
					res += r.dp(i+int(damagedCount)+1, j+1, c)
				} else {
					res += r.dp(i+int(damagedCount), j+1, c)
				}
			}
		}
	}
	c[i][j] = res
	return res
}

func (r *Record) countArrangements() int {
	var c [][]int
	for i := 0; i < len(r.record); i++ {
		c = append(c, make([]int, len(r.groups)+1))
		for j := 0; j < len(r.groups)+1; j++ {
			c[i][j] = -1
		}
	}

	return r.dp(0, 0, c)
}

func (r *Record) unfold(times int) {
	initRec := r.record
	initGrps := r.groups

	for i := 0; i < times-1; i++ {
		r.record = append(r.record, '?')
		r.record = append(r.record, initRec...)
		r.groups = append(r.groups, initGrps...)
	}
}

func playPart0(fileName string) int {
	input := readFile(fileName)
	records := parseForPart0(input)
	result := 0
	for i, r := range records {
		newVal := r.countArrangements()
		result += newVal
		if i%100 == 0 {
			fmt.Printf("Record %d: %+v\nCount: %d\nResult: %d\n\n", i, *r, newVal, result)
		}
	}
	return result
}

func playPart1(fileName string) int {
	input := readFile(fileName)
	records := parseForPart0(input)
	result := 0
	for i, r := range records {
		r.unfold(5)
		newVal := r.countArrangements()
		result += newVal
		if i%100 == 0 {
			fmt.Printf("Record %d: %+v\nCount: %d\nResult: %d\n\n", i, *r, newVal, result)
		}
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

func main() {
	retVal := playPart0("test0.txt")
	fmt.Println(retVal)
	if retVal != 21 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart0("input.txt")
	fmt.Println(retVal)
	if retVal != 7169 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 525152 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != 1738259948652 {
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
