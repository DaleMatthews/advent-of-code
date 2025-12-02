package day02

import (
	"strconv"
	"strings"
)

func solve(input string, isInvalid func(string) bool) int {
	ranges := parseInput(input)
	invalids := make([]int, 0)
	for _, r := range ranges {
		low, high := r[0], r[1]

		for id := low; id <= high; id++ {
			if isInvalid(strconv.Itoa(id)) {
				invalids = append(invalids, id)
			}
		}
	}
	total := 0
	for _, num := range invalids {
		total += num
	}
	return total
}

func SolvePuzzle1(input string) int {
	return solve(input, repeatsTwice)
}

func SolvePuzzle2(input string) int {
	return solve(input, repeatsAtLeastTwice)
}

func repeatsTwice(input string) bool {
	length := len(input)
	if length%2 != 0 {
		return false
	}
	return input[0:length/2] == input[length/2:length]
}

// input has at most 10 digit numbers, so we only need
// factors of 2-10
var lengthToRepeatableLengths = map[int][]int{
	1:  {1},
	2:  {1},
	3:  {1},
	4:  {1, 2},
	5:  {1},
	6:  {1, 2, 3},
	7:  {1},
	8:  {1, 2, 4},
	9:  {1, 3},
	10: {1, 2, 5},
}

func repeatsAtLeastTwice(s string) bool {
	// return strings.Contains((s + s)[1:len(s+s)-1], s)
	length := len(s)
	if length == 1 {
		return false
	}
	repeatableLengths := lengthToRepeatableLengths[length]
	for _, step := range repeatableLengths {
		start := 0
		str := s[start:step]
		for start != length {
			if str != s[start:start+step] {
				break
			}
			start += step
		}
		if start == length {
			// we made it all the way through
			return true
		}
	}
	return false
}

func parseInput(input string) [][2]int {
	groups := strings.Split(input, ",")
	ranges := make([][2]int, 0)

	for _, group := range groups {
		strNums := strings.Split(group, "-")
		lhs, _ := strconv.Atoi(strNums[0])
		rhs, _ := strconv.Atoi(strNums[1])
		ranges = append(ranges, [2]int{lhs, rhs})
	}
	return ranges
}
