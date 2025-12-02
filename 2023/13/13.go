package main

import (
	"fmt"
	"strings"
)

func main() {
	input := input1
	puzzles := parseInput(input)
	sum := 0
	for _, p := range puzzles {
		puzzleCount, isVertical := findSymmetry(p, false)
		if isVertical {
			sum += puzzleCount
		} else {
			sum += puzzleCount * 100
		}
	}
	fmt.Println("Part 1 sum: ", sum)
	sum = 0
	for _, p := range puzzles {
		puzzleCount, isVertical := findSymmetry(p, true)
		if isVertical {
			sum += puzzleCount
		} else {
			sum += puzzleCount * 100
		}
	}
	fmt.Println("Part 2 sum: ", sum)
}

func findSymmetry(puzzle []string, part2 bool) (int, bool) {
	// verticals
	allVerticalReflections := make([]int, len(puzzle[0]))
	for _, str := range puzzle {
		updateReflectionPoints(str, &allVerticalReflections)
	}
	mirrorSize := len(puzzle)
	if part2 {
		mirrorSize = len(puzzle) - 1
	}
	for i, reflection := range allVerticalReflections {
		if reflection == mirrorSize {
			return i, true
		}
	}

	// horizontals
	allHorizontalReflections := make([]int, len(puzzle))
	for i := 0; i < len(puzzle[0]); i++ {
		str := ``
		for j := range puzzle {
			str += string(puzzle[j][i])
		}
		updateReflectionPoints(str, &allHorizontalReflections)
	}

	mirrorSize = len(puzzle[0])
	if part2 {
		mirrorSize = len(puzzle[0]) - 1
	}
	for i, reflection := range allHorizontalReflections {
		if reflection == mirrorSize {
			return i, false
		}
	}
	panic("we never found a reflection point")
}

func updateReflectionPoints(str string, reflections *[]int) {
	for i := 1; i < len(str); i++ {
		line1 := str[:i]
		line2 := str[i:]
		if len(line1) > len(line2) {
			line1 = line1[len(line1)-len(line2):]
		} else {
			line2 = line2[:len(line1)]
		}
		symmetrical := true
		for j := 0; j < len(line1); j++ {
			if line1[j] != line2[len(line2)-1-j] {
				symmetrical = false
				break
			}
		}
		if symmetrical {
			(*reflections)[i] += 1
		}
	}
}

func parseInput(input string) [][]string {
	sections := strings.Split(input, "\n\n")
	puzzles := make([][]string, len(sections))
	for i, section := range sections {
		puzzles[i] = strings.Split(section, "\n")
	}
	return puzzles
}

var input1 = `fixme`

var input2 = `fixme`
