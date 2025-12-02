package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var inputList1 = []string{"fixme"}

var inputList2 = []string{"fixme"}

func main() {
	var inputList = inputList1
	winningCards, myCards := parseCards(inputList)
	// fmt.Println(winningCards)
	// fmt.Println(myCards)

	// get match count
	cardMatchCounts := make([]int, len(inputList))
	for i, winningCard := range winningCards {
		intersection := getIntersection(winningCard, myCards[i])
		cardMatchCounts[i] = len(intersection)
	}

	// create a copy count for each scorecard and initialize to 1 copy
	totalCopies := make([]int, len(inputList))
	for i := range totalCopies {
		totalCopies[i] = 1
	}
	fmt.Println("cardMatchCounts", cardMatchCounts)

	// increment copy count as needed
	for i, count := range cardMatchCounts {
		upperBound := int(math.Min(float64(len(totalCopies)), float64(i+1+count)))
		slice := totalCopies[i+1 : upperBound]
		multiplier := totalCopies[i]
		// fmt.Println("index", i, "start", i+1, "end", upperBound)
		for j := range slice {
			slice[j] += 1 * multiplier
		}
		// fmt.Println("totalCopies", totalCopies, "\n")
	}
	fmt.Println("totalCopies", totalCopies)

	scratchCardCount := 0
	for _, copies := range totalCopies {
		scratchCardCount += copies
	}
	fmt.Println("scratchCardCount: ", scratchCardCount)
}

type set = map[int]bool

func parseCards(inputList []string) ([]set, []set) {
	winningCards := make([]set, len(inputList))
	myCards := make([]set, len(inputList))

	for i, input := range inputList {
		// parse
		numbers := strings.Split(input, ": ")
		cards := strings.Split(numbers[1], " | ")
		winningNumberStrings := strings.Fields(cards[0])
		myNumberStrings := strings.Fields(cards[1])

		// init and add to sets
		winningCards[i] = make(set)
		for _, winningNumberString := range winningNumberStrings {
			parsedNumber, _ := strconv.Atoi(winningNumberString)
			winningCards[i][parsedNumber] = true
		}
		myCards[i] = make(set)
		for _, myNumberString := range myNumberStrings {
			parsedNumber, _ := strconv.Atoi(myNumberString)
			myCards[i][parsedNumber] = true
		}
	}

	return winningCards, myCards
}

func getIntersection(s1 set, s2 set) set {
	s := make(set)
	for elem := range s1 {
		_, hasElem := s2[elem]
		if hasElem {
			s[elem] = true
		}
	}
	return s
}
