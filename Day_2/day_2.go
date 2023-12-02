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
			//fmt.Println(playStrings, len(playStrings), plays)
			currPlay := make(map[string]int)
			for idx := 0; idx+1 <= len(playStrings); idx = idx + 2 {
				numBalls, _ := strconv.Atoi(playStrings[idx])
				currPlay[strings.TrimSpace(playStrings[idx+1])] = numBalls
				//fmt.Println(idx, numBalls, currPlay, plays)
			}
			newGame.plays[len(newGame.plays)] = currPlay
			fmt.Println("Play: ", currPlay, "Plays: ", newGame.plays)
		}
		retVal = append(retVal, newGame)
		fmt.Println("Finished!!   Line: ", idx+1, "Game: ", newGame)
	}
	return retVal
}

func parseForPart2(fileName string) [][]int {
	return nil
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
	fmt.Println(games)
	reference := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	fmt.Println("Reference: ", reference)
	result := 0
	for _, game := range games {
		playFound := false
		fmt.Println("Game: ", game, game.plays[0], len(game.plays[0]))
		for _, play := range game.plays {
			fmt.Println("Play: ", play)
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

	return 0
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
