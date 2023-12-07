package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	id        int
	time      int
	distance  int
	waysToWin int
}

func (race *Race) determineWinOptions() {
	winOpts := 0
	for currHold := 0; currHold < race.time; currHold++ {
		if currHold*(race.time-currHold) > race.distance {
			winOpts++
		}
	}
	race.waysToWin = winOpts
}

func main() {
	retVal := playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 288 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != 449550 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart2("test0.txt")
	fmt.Println(retVal)
	if retVal != 71503 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = playPart2("input.txt")
	fmt.Println(retVal)
	if retVal != 28360140 {
		panic("Part 1 failed")
	}
	fmt.Println("Part 1 passed")
}

func playPart1(fileName string) int {
	input := readFile(fileName)
	races := parseForPart1(input)
	result := 1
	for _, race := range races {
		race.determineWinOptions()
		result *= race.waysToWin
	}
	return result
}

func playPart2(fileName string) int {
	input := readFile(fileName)
	races := parseForPart2(input)
	result := 1
	for _, race := range races {
		race.determineWinOptions()
		result *= race.waysToWin
	}
	return result
}

func readFile(fileName string) []string {
	bytes, _ := os.ReadFile(fileName)
	input := string(bytes)
	splitInput := strings.Split(input, "\n")
	return splitInput
}

func parseForPart1(input []string) []Race {
	timeStrings := strings.Fields(input[0])
	distanceStrings := strings.Fields(input[1])
	races := make([]Race, 0)
	for idx := 1; idx < len(timeStrings); idx++ {
		timeVal, _ := strconv.Atoi(timeStrings[idx])
		distVal, _ := strconv.Atoi(distanceStrings[idx])
		races = append(races, Race{time: timeVal, distance: distVal, id: idx})
	}
	return races
}

func parseForPart2(input []string) []Race {
	input[0] = strings.TrimSpace(input[0])
	input[1] = strings.TrimSpace(input[1])
	strippedTimes := strings.ReplaceAll(input[0], " ", "")
	strippedDistances := strings.ReplaceAll(input[1], " ", "")

	time, _ := strconv.Atoi(strings.Split(strippedTimes, ":")[1])
	dist, _ := strconv.Atoi(strings.Split(strippedDistances, ":")[1])
	races := make([]Race, 0)
	races = append(races, Race{id: 0, time: time, distance: dist})
	return races
}
