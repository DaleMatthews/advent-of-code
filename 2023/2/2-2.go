package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Set struct {
	red   *int
	green *int
	blue  *int
}
type Game = []Set

var inputList = []string{"fixme"}

// Game 1: 18 red, 8 green, 7 blue; 15 red, 4 blue, 1 green; 2 green, 17 red, 6 blue; 5 green, 1 blue, 11 red; 18 red, 1 green, 14 blue; 8 blue
func main() {
	sum := 0
	for _, input := range inputList {
		gameString := strings.Split(input, ": ")[1]
		game := getGame(gameString)
		maxSet := getMaxSet(&game)
		sum += *maxSet.red * *maxSet.green * *maxSet.blue
	}
	fmt.Println("total sum: ", sum)
}

func getGame(gameStr string) Game {
	setStrings := strings.Split(gameStr, "; ")
	sets := make([]Set, len(setStrings))
	for i, setStr := range setStrings {
		updateSet(&sets[i], setStr)
	}
	return sets
}

var reR = regexp.MustCompile(`([\d]*) red`)
var reG = regexp.MustCompile(`([\d]*) green`)
var reB = regexp.MustCompile(`([\d]*) blue`)

func updateSet(set *Set, setStr string) {
	matchR := reR.FindStringSubmatch(setStr)
	matchG := reG.FindStringSubmatch(setStr)
	matchB := reB.FindStringSubmatch(setStr)
	if len(matchR) >= 2 {
		red, err := strconv.Atoi(matchR[1])
		if err == nil {
			set.red = &red
		}
	}
	if len(matchG) >= 2 {
		green, err := strconv.Atoi(matchG[1])
		if err == nil {
			set.green = &green
		}
	}
	if len(matchB) >= 2 {
		blue, err := strconv.Atoi(matchB[1])
		if err == nil {
			set.blue = &blue
		}
	}
}

func getMaxSet(game *Game) Set {
	var red, green, blue int
	maxSet := Set{&red, &green, &blue}
	for _, set := range *game {
		if set.red != nil && *set.red > *maxSet.red {
			maxSet.red = set.red
		}
		if set.green != nil && *set.green > *maxSet.green {
			maxSet.green = set.green
		}
		if set.blue != nil && *set.blue > *maxSet.blue {
			maxSet.blue = set.blue
		}
	}
	return maxSet
}
