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

	sum := 0
	for i, winningCard := range winningCards {
		intersection := getIntersection(winningCard, myCards[i])
		sum += int(math.Pow(2, float64(len(intersection)-1)))
	}
	fmt.Println("Sum: ", sum)
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
