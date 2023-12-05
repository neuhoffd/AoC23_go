package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Map struct {
	name                                  string
	destinationStarts, sourceStarts, rngs []int
	sourceNums, destinationNums           [][]int
}

type State struct {
	seeds []int
	maps  []Map
}

func readFile(fileName string) []string {
	bytes, _ := os.ReadFile(fileName)
	input := string(bytes)
	splitInput := strings.Split(input, "\n")
	return splitInput
}

func parseForPart1(input []string) *State {
	maps := []Map{}
	state := &State{}
	var currMap *Map
	for _, line := range input {
		splits := strings.Fields(strings.TrimSpace(line))
		if len(splits) <= 1 {
			if currMap != nil {
				maps = append(maps, *currMap)
			}
			continue
		}
		if splits[0] == "seeds:" {
			for _, split := range splits[1:] {
				val, _ := strconv.Atoi(split)
				state.seeds = append(state.seeds, val)
			}
			continue
		}
		_, err := strconv.Atoi(splits[0])
		if err != nil {
			currMap = &Map{name: splits[0]}
		} else {
			val, _ := strconv.Atoi(splits[0])
			currMap.destinationStarts = append(currMap.destinationStarts, val)
			val, _ = strconv.Atoi(splits[1])
			currMap.sourceStarts = append(currMap.sourceStarts, val)
			val, _ = strconv.Atoi(splits[2])
			currMap.rngs = append(currMap.rngs, val)
		}
	}
	if currMap != nil {
		maps = append(maps, *currMap)
	}
	state.maps = maps
	return state
}

func playPart1(fileName string) int {
	lines := readFile(fileName)
	state := parseForPart1(lines)
	locations := []int{}
	currPos := 0
	for _, currSeed := range state.seeds {
		fmt.Println("Seed: ", currSeed)
		currPos = currSeed
		for _, currMap := range state.maps {
			destinationFound := false
			for idx, currSourceStart := range currMap.sourceStarts {
				if currPos >= currSourceStart && currPos < currSourceStart+currMap.rngs[idx] {
					currOffset := currPos - currSourceStart
					currPos = currMap.destinationStarts[idx] + currOffset
					destinationFound = true
				}
				if destinationFound {
					break
				}
			}
			fmt.Printf("Map: %s Position: %d\n", currMap.name, currPos)
		}
		locations = append(locations, currPos)
	}
	sort.Ints(locations)
	return locations[0]
}

func playPart2(fileName string) int {
	lines := readFile(fileName)
	state := parseForPart1(lines)
	currPos := 0
	retVal := 0
	for currLocation := 0; true; currLocation++ {
		if currLocation%10000000 == 0 {
			fmt.Println("Location: ", currLocation)
		}
		currPos = currLocation
		for mapIdx := len(state.maps) - 1; mapIdx >= 0; mapIdx-- {
			currMap := state.maps[mapIdx]
			sourceFound := false
			for idx, currDestinationStart := range currMap.destinationStarts {
				if currPos >= currDestinationStart && currPos < currDestinationStart+currMap.rngs[idx] {
					currOffset := currPos - currDestinationStart
					currPos = currMap.sourceStarts[idx] + currOffset
					sourceFound = true
				}
				if sourceFound {
					break
				}
			}
		}
		seedFound := false
		for idx := 0; idx < len(state.seeds)-1; idx += 2 {
			if currPos >= state.seeds[idx] && currPos < state.seeds[idx]+state.seeds[idx+1] {
				seedFound = true
				break
			}
		}
		if seedFound {
			retVal = currLocation
			break
		}
	}
	return retVal
}

func main() {
	retVal := playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 35 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != 251346198 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart2("test0.txt")
	fmt.Println(retVal)
	if retVal != 46 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = playPart2("input.txt")
	fmt.Println(retVal)
	if retVal != 72263011 {
		panic("Part 1 failed")
	}
	fmt.Println("Part 1 passed")
}
