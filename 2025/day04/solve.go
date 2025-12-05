package day04

import "strings"

func SolvePuzzle1(input string) int {
	rows := parseInput(input)
	accessibleRolls := 0
	for i := range rows {
		for j := range rows[i] {
			if rows[i][j] == "@" && numRollNeighbors(rows, i, j) < 4 {
				accessibleRolls++
			}
		}
	}
	return accessibleRolls
}

func SolvePuzzle2(input string) int {
	rows := parseInput(input)
	accessibleRolls := 0
	someRemoved := true
	for someRemoved {
		someRemoved = false
		for i := range rows {
			for j := range rows[i] {
				if rows[i][j] == "@" && numRollNeighbors(rows, i, j) < 4 {
					rows[i][j] = "x"
					someRemoved = true
					accessibleRolls++
				}
			}
		}
	}
	return accessibleRolls
}

// -1,-1  -1,0  -1,1
//  0,-1   0,0   0,1
//  1,-1   1,0   1,1
func numRollNeighbors(p [][]string, i, j int) int {
	height := len(p)
	width := len(p[0])
	numRolls := 0
	if i-1 >= 0 && j-1 >= 0 && p[i-1][j-1] == "@" {
		numRolls++
	}
	if i-1 >= 0 && p[i-1][j] == "@" {
		numRolls++
	}
	if i-1 >= 0 && j+1 < width && p[i-1][j+1] == "@" {
		numRolls++
	}
	if j-1 >= 0 && p[i][j-1] == "@" {
		numRolls++
	}
	if j+1 < width && p[i][j+1] == "@" {
		numRolls++
	}
	if i+1 < height && j-1 >= 0 && p[i+1][j-1] == "@" {
		numRolls++
	}
	if i+1 < height && p[i+1][j] == "@" {
		numRolls++
	}
	if i+1 < height && j+1 < width && p[i+1][j+1] == "@" {
		numRolls++
	}
	return numRolls
}

func parseInput(input string) [][]string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	board := make([][]string, len(lines))

	for i, line := range lines {
		board[i] = strings.Split(strings.TrimSpace(line), "")
	}
	return board
}
