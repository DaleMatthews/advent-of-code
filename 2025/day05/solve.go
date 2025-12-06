package day05

import (
	"slices"
	"spissable/advent-of-go-template/utils"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) int {
	ranges, ids := parseInput(input)
	numFresh := 0
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				numFresh++
				break
			}
		}
	}
	return numFresh
}

func SolvePuzzle2(input string) int {
	ranges, _ := parseInput(input)
	numFresh := 0
	for _, r := range ranges {
		numFresh += r[1] - r[0] + 1
	}
	return numFresh
}

func parseInput(input string) ([][2]int, []int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ranges := make([][2]int, 0)
	ids := make([]int, 0)

	for _, line := range lines {
		if strings.Contains(line, "-") {
			vals := strings.Split(line, "-")
			ranges = append(ranges, [2]int{
				utils.Must(strconv.Atoi(strings.TrimSpace(vals[0]))),
				utils.Must(strconv.Atoi(strings.TrimSpace(vals[1]))),
			})
		} else if len(strings.TrimSpace(line)) > 0 {
			ids = append(ids, utils.Must(strconv.Atoi(strings.TrimSpace(line))))
		}
	}
	return combineRanges(ranges), ids
}

func combineRanges(ranges [][2]int) [][2]int {
	combined := make([][2]int, 0)
	for _, r := range ranges {
		segmented := true
		for j, c := range combined {
			lhs := inRange(c[0], r[0], r[1])
			rhs := inRange(c[1], r[0], r[1])
			if lhs || rhs {
				segmented = false
				c[0] = min(c[0], r[0])
				c[1] = max(c[1], r[1])
				combined[j] = c
				break
			}
		}
		if segmented {
			combined = append(combined, r)
		}
	}
	if len(ranges) > len(combined) {
		println("turned", len(ranges), "ranges into", len(combined), "ranges")
		// bit hacky but this works...
		slices.Reverse(combined)
		return combineRanges(combined)
	}
	return combined
}

func inRange(x, r1, r2 int) bool {
	return x >= r1 && x <= r2
}
