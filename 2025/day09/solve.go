package day09

import "strings"

func SolvePuzzle1(input string) int {
	board := parseInput(input)
	return len(board)
}

func SolvePuzzle2(input string) int {
	board := parseInput(input)
	return len(board)
}

func parseInput(input string) [][]string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	board := make([][]string, len(lines))

	for i, line := range lines {
		board[i] = strings.Split(strings.TrimSpace(line), "")
	}
	return board
}
