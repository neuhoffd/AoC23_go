package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type stepInput struct {
	w    int64
	z    int64
	vars []int64
	idx  int
}

type statePrioritySet struct {
	zMap        map[int64]int64
	cardinality int64
}

func (s *statePrioritySet) add(key int64, newValue int64) {
	s.zMap[key] = newValue
	s.cardinality++
}

func (s *statePrioritySet) update(key int64, newValue int64) {
	s.zMap[key] = newValue
}

func (s *statePrioritySet) addIfBigger(key int64, newValue int64) {
	previousValue, ok := s.zMap[key]
	if ok {
		if previousValue < newValue {
			s.update(key, newValue)
		}
	} else {
		s.add(key, newValue)
	}
}

func (s *statePrioritySet) addIfSmaller(key int64, newValue int64) {
	previousValue, ok := s.zMap[key]
	if ok {
		if previousValue > newValue {
			s.update(key, newValue)
		}
	} else {
		s.add(key, newValue)
	}
}

func step(input stepInput) int64 {
	if len(input.vars) != 10 {
		panic("Vars not of correct length")
	}
	z := input.z
	w := input.w
	var x int64 = 0
	var y int64 = 0
	x *= input.vars[0]
	x += z
	x %= input.vars[1]
	z = int64(z / input.vars[2])
	x += input.vars[3]
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == input.vars[4] {
		x = 1
	} else {
		x = 0
	}
	y *= input.vars[5]
	y += input.vars[6]
	y *= x
	y += input.vars[7]
	z *= y
	y *= input.vars[8]
	y += w
	y += input.vars[9]
	y *= x
	z += y

	return z
}

func generateNewStates(vars *[140]int64, digit int64, idx int, minimum bool, prevStates *statePrioritySet, newStates *statePrioritySet) {
	for oldZ, oldModelNumber := range prevStates.zMap {
		nextStep := stepInput{w: int64(digit), z: oldZ, vars: vars[idx*10 : idx*10+10], idx: idx}
		newZ := step(nextStep)
		if minimum {
			newStates.addIfSmaller(newZ, oldModelNumber*10+digit)
		} else {
			newStates.addIfBigger(newZ, oldModelNumber*10+digit)
		}
	}
}

func searchModelNumber(vars [140]int64, minimum bool) int64 {
	prevStates := &statePrioritySet{zMap: make(map[int64]int64)}
	start := time.Now()
	prevStates.zMap[int64(0)] = 0
	for digitIdx := 0; digitIdx < 14; digitIdx++ {
		passStart := time.Now()
		fmt.Printf("Starting pass %d \n", digitIdx)
		newStates := statePrioritySet{zMap: make(map[int64]int64)}
		for digit := int64(9); digit > 0; digit-- {
			generateNewStates(&vars, digit, digitIdx, minimum, prevStates, &newStates)
		}
		prevStates = &newStates
		fmt.Printf("Pass %d finished after %.2f seconds. Cardinality : %d\n", digitIdx, time.Since(passStart).Seconds(), prevStates.cardinality)
	}
	fmt.Printf("Model number search done after %.2f seconds\n", time.Since(start).Seconds())
	return prevStates.zMap[0]
}

func parseVariables(fileName string) [140]int64 {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR!", err)
		panic(fmt.Sprintf("Could not open file %s", fileName))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var vars [140]int64
	var curr int64
	var line string
	var garbage1, garbage2 string
	counter := 0
	for scanner.Scan() {
		line = scanner.Text()
		_, err := fmt.Sscanf(line, "%s %s %d", &garbage1, &garbage2, &curr)
		if err == nil {
			vars[counter] = curr
			counter++
		}
	}
	return vars
}

func main() {
	vars := parseVariables("input.txt")
	fmt.Println(vars)

	biggest := searchModelNumber(vars, false)
	fmt.Println("Biggest No: ", biggest)

	smallest := searchModelNumber(vars, true)
	fmt.Println("Smallest No: ", smallest)
}
