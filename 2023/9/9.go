package main

import (
	"1/util"
	"fmt"
	"strings"
)

var input1 string = `fixme`

var input2 string = `fixme`

func main() {
	input := input1
	histories := parseInput(input)
	nextSum := 0
	prevSum := 0
	for i, history := range histories {
		next, deriveCount := nextValue(history, 0)
		nextSum += next
		fmt.Println("Line ", i+1, " took ", deriveCount, " derivatives")
		prevSum += prevValue(history)
	}
	fmt.Println("nextSum: ", nextSum)
	fmt.Println("prevSum: ", prevSum)
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	histories := make([][]int, len(lines))
	for i, line := range lines {
		histories[i] = util.StringsToInts(strings.Fields(line))
	}
	return histories
}

func nextValue(history []int, count int) (int, int) {
	// 0   3   6   9   12   15   B
	//   3   3   3   3   3   A
	//     0   0   0   0   0
	lastIndex := len(history) - 1
	if history[lastIndex] == 0 {
		return 0, count
	}
	derivedHistory := make([]int, lastIndex)
	for i, datum := range history[1:] {
		derivedHistory[i] = datum - history[i]
	}
	next, count := nextValue(derivedHistory, count+1)
	return history[lastIndex] + next, count
}

func prevValue(history []int) int {
	// B   0   3   6   9  12  15
	//   A   3   3   3   3   3
	//     0   0   0   0   0
	lastIndex := len(history) - 1
	if history[lastIndex] == 0 {
		return 0
	}
	derivedHistory := make([]int, lastIndex)
	for i, datum := range history[1:] {
		derivedHistory[i] = datum - history[i]
	}
	return history[0] - prevValue(derivedHistory)
}
