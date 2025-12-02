package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var input1 string = `fixme`

var input2 string = `fixme`

var input3 string = `fixme`

var cardToValuePart1 = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}
var cardToValuePart2 = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 1,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type Hand struct {
	cards      string
	totalValue int
	handType   int
	bid        int
}

func main() {
	input := input1
	hands := parseInput(input, true)
	sortHands(hands)
	// for _, hand := range hands {
	// 	fmt.Println(hand.handType, hand.cards)
	// }
	fmt.Println(bidRankSum(hands))
}

func parseInput(input string, part2 bool) []Hand {
	cardToValue := cardToValuePart1
	if part2 {
		cardToValue = cardToValuePart2
	}
	lines := strings.Split(input, "\n")
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		cards := fields[0]
		bid, _ := strconv.Atoi(fields[1])
		cardValues := make([]int, 5)
		for j, card := range cards {
			cardValues[j] = cardToValue[card]
		}
		var handType int
		if part2 {
			handType = cardsToHandTypePart2(cards)
		} else {
			handType = cardsToHandType(cards)
		}
		totalValue := handType + cardValues[0]*100000000 + cardValues[1]*1000000 + cardValues[2]*10000 + cardValues[3]*100 + cardValues[4]
		hands[i] = Hand{
			cards:      cards,
			totalValue: totalValue,
			handType:   handType,
			bid:        bid,
		}
	}
	return hands
}

const (
	HighCard     int = 10000000000
	OnePair          = 20000000000
	TwoPair          = 30000000000
	ThreeOfAKind     = 40000000000
	FullHouse        = 50000000000
	FourOfAKind      = 60000000000
	FiveOfAKind      = 70000000000
)

func cardsToHandType(cards string) int {
	var cardCount = map[rune]int{
		'2': 0,
		'3': 0,
		'4': 0,
		'5': 0,
		'6': 0,
		'7': 0,
		'8': 0,
		'9': 0,
		'T': 0,
		'J': 0,
		'Q': 0,
		'K': 0,
		'A': 0,
	}
	for _, card := range cards {
		cardCount[card]++
	}
	pairFound := false
	threeFound := false
	for _, count := range cardCount {
		switch {
		case count == 5:
			return FiveOfAKind
		case count == 4:
			return FourOfAKind
		case count == 3:
			if pairFound {
				return FullHouse
			}
			threeFound = true
		case count == 2:
			if threeFound {
				return FullHouse
			} else if pairFound {
				return TwoPair
			}
			pairFound = true
		}
	}
	if threeFound {
		return ThreeOfAKind
	}
	if pairFound {
		return OnePair
	}
	return HighCard
}

func cardsToHandTypePart2(cards string) int {
	var cardCount = map[rune]int{
		'2': 0,
		'3': 0,
		'4': 0,
		'5': 0,
		'6': 0,
		'7': 0,
		'8': 0,
		'9': 0,
		'T': 0,
		'Q': 0,
		'K': 0,
		'A': 0,
	}
	jokerCount := 0
	for _, card := range cards {
		if card == 'J' {
			jokerCount++
		} else {
			cardCount[card]++
		}
	}
	if jokerCount > 3 {
		return FiveOfAKind
	}
	pairCount := 0
	threeFound := false
	for _, count := range cardCount {
		switch {
		case count == 5:
			return FiveOfAKind
		case count == 4:
			if jokerCount == 1 {
				return FiveOfAKind
			}
			return FourOfAKind
		case count == 3:
			threeFound = true
		case count == 2:
			pairCount++
		}
	}
	// AT THIS POINT, JOKERCOUNT IS ONLY 0-3
	if threeFound {
		if pairCount == 1 {
			return FullHouse
		}
		if jokerCount == 2 {
			return FiveOfAKind
		}
		if jokerCount == 1 {
			return FourOfAKind
		}
		return ThreeOfAKind
	}
	if pairCount == 2 {
		if jokerCount == 1 {
			return FullHouse
		}
		return TwoPair
	}
	if pairCount == 1 {
		if jokerCount == 0 {
			return OnePair
		} else if jokerCount == 1 {
			return ThreeOfAKind
		} else if jokerCount == 2 {
			return FourOfAKind
		} else if jokerCount == 3 {
			return FiveOfAKind
		}
	}
	if jokerCount == 3 {
		return FourOfAKind
	} else if jokerCount == 2 {
		return ThreeOfAKind
	} else if jokerCount == 1 {
		return OnePair
	}
	return HighCard
}

func sortHands(hands []Hand) {
	cmp := func(a, b Hand) int {
		return cmp.Compare(a.totalValue, b.totalValue)
	}
	slices.SortFunc(hands, cmp)
}

func bidRankSum(sortedHands []Hand) int {
	sum := 0
	for i, hand := range sortedHands {
		sum += (i + 1) * hand.bid
	}
	return sum
}
