package day03

import (
	"slices"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) int {
	banks := parseInput(input)
	totalJoltage := 0

	for _, bank := range banks {
		firstMax := slices.Max(bank[:len(bank)-1])
		i := slices.Index(bank, firstMax)
		secondMax := slices.Max(bank[i+1:])
		joltage, _ := strconv.Atoi(string(rune(firstMax+'0')) + string(rune(secondMax+'0')))
		println("joltage", firstMax, secondMax, joltage)
		totalJoltage += joltage
	}

	return totalJoltage
}

func SolvePuzzle2(input string) int {
	banks := parseInput(input)
	totalJoltage := 0

	for _, bank := range banks {
		start := 0
		joltageStr := ""
		for i := 0; i < 12; i++ {
			max := slices.Max(bank[start : len(bank)-(11-i)])
			joltageStr += strconv.Itoa(max)
			start += slices.Index(bank[start:], max) + 1
		}
		joltage, _ := strconv.Atoi(joltageStr)
		println("joltage", joltage)
		totalJoltage += joltage
	}

	return totalJoltage
}

func parseInput(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	banks := make([][]int, len(lines))

	for i, line := range lines {
		line = strings.TrimSpace(line)
		banks[i] = make([]int, len(line))
		for j, char := range line {
			banks[i][j] = int(char - '0')
		}
	}
	return banks
}
