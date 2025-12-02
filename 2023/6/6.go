package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var input1 string = `fixme`

var input2 string = `fixme`

// part 2 real input
var input3 string = `fixme`

// part 2 mock input
var input4 string = `fixme`

func main() {
	input := input3
	times, records := parseInput(input)
	total := 1
	for i, allowedTime := range times {
		currentRecord := records[i]
		upperZero := (allowedTime + math.Sqrt(math.Pow(allowedTime, 2)-4*currentRecord)) / 2
		lowerZero := (allowedTime - math.Sqrt(math.Pow(allowedTime, 2)-4*currentRecord)) / 2
		upperBound := math.Floor(upperZero - 0.0000001)
		lowerBound := math.Ceil(lowerZero + 0.0000001)
		possibleWins := upperBound - lowerBound + 1
		fmt.Println(lowerBound, " ... ", upperBound)
		// fmt.Println("possible wins", possibleWins)
		total *= int(possibleWins)
	}
	fmt.Println("possible wins compounded", total)
}

func parseInput(input string) ([]float64, []float64) {
	lines := strings.Split(input, "\n")
	times := strings.Fields(lines[0])[1:]
	records := strings.Fields(lines[1])[1:]
	return stringsToFloats(times), stringsToFloats(records)
}

func stringsToFloats(strs []string) []float64 {
	float64s := make([]float64, len(strs))
	for i, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		float64s[i] = float64(num)
	}
	return float64s
}

/**
Time:      7
Distance:  9

T=allowed time=7
H=held time
R=record=9
Find values of H such that (T-H)*H > R
T-H > R/H
(7-H)*H > 9
7H-H^2 > 9
H^2 - TH + R <= 0

It's a quadratic formula, how bout that
a=1
b=-T
c=R

H = (T +- sqrt(T^2 - 4R)) / 2
H = (7 +- sqrt(7^2 - 4*9)) / 2
H = (7 +- sqrt(49 - 36)) / 2
H = (7 +- sqrt(13)) / 2
*/
