package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	winning, own []int
}

func parseForPart1(input []string) []Card {
	cards := []Card{}

	for _, line := range input {
		parsingWinning := true
		actCard := Card{}
		r := regexp.MustCompile(`\||\s+`)
		splits := r.Split(line, -1)
		fmt.Println(splits)
		for _, part := range splits[2:] {
			if part == "" {
				parsingWinning = false
				continue
			}
			val, _ := strconv.Atoi(part)
			if parsingWinning {
				actCard.winning = append(actCard.winning, val)
			} else {
				actCard.own = append(actCard.own, val)
			}
		}
		cards = append(cards, actCard)
	}
	return cards
}

func readFile(fileName string) []string {
	bytes, _ := os.ReadFile(fileName)
	input := string(bytes)
	splitInput := strings.Split(input, "\n")
	return splitInput
}

func playPart1(fileName string) int {
	input := readFile(fileName)
	cards := parseForPart1(input)
	result := 0
	for idx, card := range cards {
		count := 0
		for _, winner := range card.winning {
			if slices.Contains(card.own, winner) {
				count++
			}
		}
		score := 0
		if count > 0 {
			score = int(math.Pow(2, float64(count-1)))
		}
		result += score
		fmt.Printf("Card %d, Count: %d Score: %d \n", idx, count, score)
	}

	return result
}

func playPart2(fileName string) int {

	return 0
}

func main() {
	retVal := playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 13 {
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
