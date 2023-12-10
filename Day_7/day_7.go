package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var remap map[string]int

func initRemap() {
	remap = map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
	}
}

type Hand struct {
	id         int
	cards      []int
	cardCounts map[int]int
	bid        int
	tp         int
}

func (hand Hand) aKind(howMany int) bool {
	for _, v := range hand.cardCounts {
		if v == howMany {
			return true
		}
	}
	return false
}

func (hand Hand) fullHouse() bool {
	fh2 := false
	fh3 := false
	for _, v := range hand.cardCounts {
		fh2 = (v == 2 || fh2)
		fh3 = (v == 3 || fh3)
	}
	return fh2 && fh3
}

func (hand Hand) pairs(num int) bool {
	pairs := make([]bool, 0)
	for _, v := range hand.cardCounts {
		if v == 2 {
			pairs = append(pairs, true)
		}
	}
	return len(pairs) == num
}

func (hand *Hand) determineType() {
	if hand.aKind(5) {
		hand.tp = 7
		return
	}
	if hand.aKind(4) {
		hand.tp = 6
		return
	}
	if hand.fullHouse() {
		hand.tp = 5
		return
	}
	if hand.aKind(3) {
		hand.tp = 4
		return
	}
	if hand.pairs(2) {
		hand.tp = 3
		return
	}
	if hand.pairs(1) {
		hand.tp = 2
		return
	}
	hand.tp = 1
}

func determineCardTypes(hands []Hand) []Hand {
	for idx := range hands {
		hands[idx].determineType()
	}
	fmt.Printf("Hands %+v \n", hands)
	return hands
}

func tieBreak(handI, handJ Hand) bool {
	for idx := 0; idx < len(handI.cards); idx++ {
		if handI.cards[idx] == handJ.cards[idx] {
			continue
		}
		if handI.cards[idx] < handJ.cards[idx] {
			return true
		} else {
			return false
		}
	}
	panic("Tie Found")
	return false
}

func sortHands(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].tp != hands[j].tp {
			return hands[i].tp < hands[j].tp
		}
		return tieBreak(hands[i], hands[j])
	})
	return hands
}

func calcWinnings(sortedHands []Hand) int {
	ans := 0
	currType := -1
	for idx, currHand := range sortedHands {
		if currHand.tp != currType {
			currType = currHand.tp
			fmt.Printf("Now Parsing type: %d\n", currType)
		}
		fmt.Printf("Current Hand: %+v\n", currHand)
		ans += (idx + 1) * currHand.bid
		//fmt.Println(ans)
	}
	return ans
}

func main() {
	initRemap()
	retVal := playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 6440 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != -1 {
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
func readFile(fileName string) []string {
	bytes, _ := os.ReadFile(fileName)
	input := string(bytes)
	splitInput := strings.Split(input, "\n")
	return splitInput
}

func playPart1(fileName string) int {
	input := readFile(fileName)
	fmt.Println(input)
	hands := parseForPart1(input)
	//fmt.Printf("Hands: %+v \n", hands)
	hands = determineCardTypes(hands)
	//fmt.Printf("Hands: %+v \n", hands)
	hands = sortHands(hands)
	//fmt.Printf("Sorted Hands: %+v \n", hands)
	return calcWinnings(hands)
}
func parseForPart1(input []string) []Hand {
	hands := make([]Hand, 0)
	for idx, line := range input {
		splits := strings.Fields(line)
		currHand := Hand{id: idx}
		currHand.cardCounts = make(map[int]int)
		for _, card := range splits[0] {
			if unicode.IsDigit(card) {
				val, _ := strconv.Atoi(string(card))
				currHand.cards = append(currHand.cards, val)
			} else {
				currHand.cards = append(currHand.cards, remap[string(card)])
			}
			currHand.cardCounts[currHand.cards[len(currHand.cards)-1]]++
		}
		val, _ := strconv.Atoi(splits[1])
		currHand.bid = val
		hands = append(hands, currHand)
	}
	return hands
}

func playPart2(fileName string) int {

	return 0
}
