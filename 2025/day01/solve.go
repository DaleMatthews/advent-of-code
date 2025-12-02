package day01

import (
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) int {
	cur := 50
	totalZeros := 0
	for _, r := range parseInput(input) {
		cur = (cur + r) % 100
		if cur == 0 {
			totalZeros++
		}
	}
	return totalZeros
}

func SolvePuzzle2(input string) int {
	cur := 50
	totalZeros := 0
	for _, r := range parseInput(input) {
		if r == 0 {
			continue
		}
		old := cur
		cur = (cur + r) % 100
		if cur < 0 {
			cur += 100
		}

		totalZeros += Abs(r) / 100
		if cur == 0 {
			totalZeros++
		}

		// edge-cases...
		remainder := Abs(r) % 100
		if remainder > 0 && old != 0 {
			if r > 0 && old+remainder > 100 {
				totalZeros++
			} else if r < 0 && old < remainder {
				totalZeros++
			}
		}
	}
	return totalZeros
}

func parseInput(input string) []int {
	var rotations []int
	for line := range strings.Lines(input) {
		line = strings.TrimSpace(line)
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			println("Error while parsing:", err.Error())
			continue
		}
		if strings.HasPrefix(line, "R") {
			rotations = append(rotations, num)
		} else {
			rotations = append(rotations, -num)
		}
	}
	return rotations
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
