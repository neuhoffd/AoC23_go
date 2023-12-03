package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	id    int
	plays map[int]map[string]int
}

func parseForPart1(input []string) []Game {
	var retVal []Game
	for idx, line := range input {
		newGame := Game{
			id:    idx + 1,
			plays: map[int]map[string]int{},
		}
		gameStrings := strings.Split(strings.Split(line, ": ")[1], "; ")
		r := regexp.MustCompile("(, +)|( +)")
		for _, gameString := range gameStrings {
			playStrings := r.Split(gameString, -1)
			currPlay := make(map[string]int)
			for idx := 0; idx+1 <= len(playStrings); idx = idx + 2 {
				numBalls, _ := strconv.Atoi(playStrings[idx])
				currPlay[strings.TrimSpace(playStrings[idx+1])] = numBalls
			}
			newGame.plays[len(newGame.plays)] = currPlay
		}
		retVal = append(retVal, newGame)
	}
	return retVal
}

func readFile(fileName string) []string {
	bytes, _ := os.ReadFile(fileName)
	input := string(bytes)
	splitInput := strings.Split(input, "\n")
	return splitInput
}

func playPart1(fileName string) int {
	lines := readFile(fileName)
	games := parseForPart1(lines)
	reference := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	result := 0
	for _, game := range games {
		playFound := false
		for _, play := range game.plays {
			for color, balls := range reference {
				if play[color] > balls {
					playFound = true
					break
				}
			}
		}
		if !playFound {
			result = result + game.id
		}
	}
	return result
}

func playPart2(fileName string) int {
	lines := readFile(fileName)
	games := parseForPart1(lines)
	result := 0
	for _, game := range games {
		minBalls := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, play := range game.plays {
			for color, currMin := range minBalls {
				if currMin < play[color] {
					minBalls[color] = play[color]
				}
			}
		}
		result = result + minBalls["red"]*minBalls["green"]*minBalls["blue"]
	}
	return result
}

func main() {
	retVal := playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 8 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != 2101 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart2("test0.txt")
	fmt.Println(retVal)
	if retVal != 2286 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = playPart2("input.txt")
	fmt.Println(retVal)
	if retVal != 58269 {
		panic("Part 1 failed")
	}
	fmt.Println("Part 1 passed")
}
