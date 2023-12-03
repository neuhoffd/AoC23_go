package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Symbol struct {
	x               []int
	y               int
	isNumber        bool
	strVal          string
	numVal          int
	neighbors       int
	id              string
	neighborNumbers []int
}

func createSymbolGrid(symbols map[string]Symbol, input []string) [][]string {
	xSize := len(input[0])
	ySize := len(input)
	symbolGrid := make([][]string, 0)
	for y := 0; y < ySize; y++ {
		symbolGrid = append(symbolGrid, make([]string, xSize))
	}
	for _, symbol := range symbols {
		for x := symbol.x[0]; x < symbol.x[1]; x++ {
			symbolGrid[symbol.y][x] = symbol.id
		}
	}
	return symbolGrid
}

func fillNeighbors(symbols map[string]Symbol, symbolGrid [][]string) map[string]Symbol {
	xSize := len(symbolGrid[0])
	ySize := len(symbolGrid)
	retVal := make(map[string]Symbol, 0)
	for _, currSymbol := range symbols {
		maxNeighbors := 0
		neighborNumbers := make([]int, 0)
		for x := currSymbol.x[0]; x < currSymbol.x[1]; x++ {
			symbolsFound := make([]string, 0)
			symbolsFound = append(symbolsFound, currSymbol.id)
			for xOff := -1; xOff <= 1; xOff++ {
				for yOff := -1; yOff <= 1; yOff++ {
					if xOff == 0 && yOff == 0 {
						continue
					}
					xCheck := x + xOff
					yCheck := currSymbol.y + yOff
					if uint(xCheck) < uint(xSize) && uint(yCheck) < uint(ySize) {
						currId := symbolGrid[yCheck][xCheck]
						if currId == "" {
							continue
						}
						alreadyFound := slices.Contains(symbolsFound, currId)
						if !alreadyFound {
							if symbols[currId].isNumber {
								neighborNumbers = append(neighborNumbers, symbols[currId].numVal)
							}
							symbolsFound = append(symbolsFound, symbolGrid[yCheck][xCheck])
						}
					}
				}
			}
			if maxNeighbors < len(symbolsFound)-1 {
				maxNeighbors = len(symbolsFound) - 1
			}
		}
		currSymbol.neighbors = maxNeighbors
		currSymbol.neighborNumbers = neighborNumbers
		retVal[currSymbol.id] = currSymbol
	}
	return retVal
}

func getSymbols(input []string) map[string]Symbol {
	r := regexp.MustCompile(`\d+|[^.]`)
	symbols := make(map[string]Symbol, 0)
	for y, line := range input {
		symbolIndices := r.FindAllStringIndex(line, -1)
		for _, currLocation := range symbolIndices {
			currString := line[currLocation[0]:currLocation[1]]
			numVal, err := strconv.Atoi(currString)
			newSymbol := Symbol{
				x:         currLocation,
				y:         y,
				isNumber:  err == nil,
				strVal:    currString,
				numVal:    numVal,
				neighbors: 0,
				id:        fmt.Sprintf("%d%d%d", currLocation[0], currLocation[1], y),
			}
			symbols[newSymbol.id] = newSymbol
		}
	}
	return symbols
}

func readFile(fileName string) []string {
	bytes, _ := os.ReadFile(fileName)
	input := string(bytes)
	splitInput := strings.Fields(input)
	return splitInput
}

func playPart1(fileName string) int {
	input := readFile(fileName)
	symbols := getSymbols(input)
	symbolGrid := createSymbolGrid(symbols, input)
	symbols = fillNeighbors(symbols, symbolGrid)
	result := 0
	for _, symbol := range symbols {
		if symbol.isNumber && symbol.neighbors > 0 {
			result = result + symbol.numVal
		}
	}
	return result
}

func playPart2(fileName string) int {
	input := readFile(fileName)
	symbols := getSymbols(input)
	symbolGrid := createSymbolGrid(symbols, input)
	symbols = fillNeighbors(symbols, symbolGrid)
	result := 0
	for _, symbol := range symbols {
		gearValue := 0
		if symbol.strVal == "*" && len(symbol.neighborNumbers) > 1 {
			gearValue = 1
			for _, val := range symbol.neighborNumbers {
				gearValue = gearValue * val
			}
		}
		result = result + gearValue
	}
	return result
}

func main() {
	retVal := playPart1("test0.txt")
	fmt.Println(retVal)
	if retVal != 4361 {
		panic("Test 0 failed")
	}
	fmt.Println("Test 0 passed")

	retVal = playPart1("input.txt")
	fmt.Println(retVal)
	if retVal != 521515 {
		panic("Part 0 failed")
	}
	fmt.Println("Part 0 passed")

	retVal = playPart2("test0.txt")
	fmt.Println(retVal)
	if retVal != 467835 {
		panic("Test 1 failed")
	}
	fmt.Println("Test 1 passed")

	retVal = playPart2("input.txt")
	fmt.Println(retVal)
	if retVal != 69527306 {
		panic("Part 1 failed")
	}
	fmt.Println("Part 1 passed")
}
