package day07

import (
	"spissable/advent-of-go-template/utils"
	"strings"
)

func SolvePuzzle1(input string) int {
	lines, start := parseInput(input)
	numSplits := 0

	beams := map[int]bool{start: true}
	for i := 0; i < len(lines); i += 2 {
		line := lines[i]

		nextBeams := make(map[int]bool)
		for beam := range beams {
			if line[beam] == '^' {
				numSplits++
				if beam > 0 {
					nextBeams[beam-1] = true
				}
				if beam+1 < len(line) {
					nextBeams[beam+1] = true
				}
			} else {
				nextBeams[beam] = true
			}
		}
		beams = nextBeams
	}

	return numSplits
}

func SolvePuzzle2(input string) int {
	lines, start := parseInput(input)
	cache := make(map[utils.Coord2D]int)
	return 1 + getNumNewTimelines(lines, 0, start, cache)
}

func getNumNewTimelines(lines []string, i int, beam int, cache map[utils.Coord2D]int) int {
	if i >= len(lines) {
		return 0
	}
	key := utils.Coord2D{X: i, Y: beam}
	if val, exists := cache[key]; exists {
		return val
	}

	line := lines[i]
	newTimelines := 0
	if line[beam] == '^' {
		if beam > 0 {
			newTimelines += getNumNewTimelines(lines, i+2, beam-1, cache)
		}
		if beam+1 < len(line) {
			newTimelines += getNumNewTimelines(lines, i+2, beam+1, cache)
		}
		if beam > 0 && beam+1 < len(line) {
			newTimelines++
		}
	} else {
		newTimelines += getNumNewTimelines(lines, i+2, beam, cache)
	}
	cache[key] = newTimelines
	return newTimelines
}

func parseInput(input string) ([]string, int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	start := strings.IndexByte(lines[0], 'S')
	return lines[2:], start
}
