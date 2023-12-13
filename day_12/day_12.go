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

func (r *Record) recurse(recordIdx, groupIdx int, dpCache [][]int) int {
	if recordIdx >= len(r.record) {
		if groupIdx < len(r.groups) {
			return 0
		}
		return 1
	}
	if dpCache[recordIdx][groupIdx] != -1 {
		return dpCache[recordIdx][groupIdx]
	}
	res := 0
	if r.record[recordIdx] == '.' {
		res += r.recurse(recordIdx+1, groupIdx, dpCache)
	} else {
		if r.record[recordIdx] == '?' {
			res += r.recurse(recordIdx+1, groupIdx, dpCache)
		}
		if groupIdx < len(r.groups) {
			damagedCount := byte(0)
			for lookAheadIdx := recordIdx; lookAheadIdx < len(r.record); lookAheadIdx++ {
				if r.record[lookAheadIdx] == '.' || (r.groups[groupIdx] == damagedCount && r.record[lookAheadIdx] == '?') || damagedCount > r.groups[groupIdx] {
					break
				}
				damagedCount++
			}
			if damagedCount == r.groups[groupIdx] {
				if recordIdx+int(damagedCount) < len(r.record) && r.record[recordIdx+int(damagedCount)] != '#' {
					res += r.recurse(recordIdx+int(damagedCount)+1, groupIdx+1, dpCache)
				} else {
					res += r.recurse(recordIdx+int(damagedCount), groupIdx+1, dpCache)
				}
			}
		}
	}
	dpCache[recordIdx][groupIdx] = res
	return res
}

func (r *Record) countArrangements() int {
	var dpCache [][]int
	for i := 0; i < len(r.record); i++ {
		dpCache = append(dpCache, make([]int, len(r.groups)+1))
		for j := 0; j < len(r.groups)+1; j++ {
			dpCache[i][j] = -1
		}
	}

	return r.recurse(0, 0, dpCache)
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
