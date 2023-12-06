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

func main() {
	retVal := playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 288 {
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

func playPart2(fileName string) int {

	return 0
}

func (race *Race) determineWinOptions() {
	//winDistance <= (time - holdtime)*holdtime <=> time*holdtime - holdtime^2 - winDist> 0 <=> holdtime^2 - time*holdtime + winDist < 0
	/*zero0 := int(math.Floor(((float64(race.time) - math.Sqrt(float64(race.time*race.time-4*race.distance))) / 2))) + 1
	zero1 := int(math.Floor(((float64(race.time) + math.Sqrt(float64(race.time*race.time-4*race.distance))) / 2)))

	race.waysToWin = zero1 - zero0
	fmt.Printf("Race: %d   Ways to win: %d \n", race.id, race.waysToWin)*/
	//Brute force
	winOpts := 0
	for currHold := 0; currHold < race.time; currHold++ {
		if currHold*(race.time-currHold) > race.distance {
			winOpts++
		}
	}
	race.waysToWin = winOpts
}
