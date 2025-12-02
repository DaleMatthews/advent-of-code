package main

import (
	"fmt"
	"strings"
)

func main() {
	input := input1
	board := parseInput(input)
	// part 1
	tiltNorthSouth(&board, -1, 0)
	fmt.Println("Sum: ", countRows(&board))

	// part 2
	board = parseInput(input)
	southStart := len(board) - 1
	westStart := len(board[0]) - 1

	// do the initial 500 to settle the board
	numSettled := 500
	for i := 0; i < numSettled; i++ {
		cycle(&board, &southStart, &westStart)
	}

	// copy this board state
	settledBoard := make([][]rune, len(board))
	for i := range board {
		settledBoard[i] = make([]rune, len(board[i]))
		copy(settledBoard[i], board[i])
	}

	// now find how often it repeats by checking when the board returns to the settled state
	sums := make([]int, 100)
	repeatEvery := -1
	for i := 0; i < 100; i++ {
		cycle(&board, &southStart, &westStart)
		sums[i] = countRows(&board)
		if i > 1 && boardsAreEqual(&settledBoard, &board) {
			repeatEvery = i + 1
			break
		}
	}
	if repeatEvery == -1 {
		fmt.Println("No pattern found")
		return
	}
	fmt.Println("Board state repeats every ", repeatEvery, " cycles")
	// calculate where in the pattern we'd be
	offset := (1000000000 - numSettled) % repeatEvery
	fmt.Println("Cycles 1000000000 is offset from this pattern by ", offset)
	if offset == 0 {
		offset = repeatEvery
	}
	fmt.Println("Cycle 1000000000 sum: ", sums[offset-1])
}

func boardsAreEqual(board1 *[][]rune, board2 *[][]rune) bool {
	for i := range *board1 {
		for j := range (*board1)[i] {
			if (*board1)[i][j] != (*board2)[i][j] {
				return false
			}
		}
	}
	return true
}

func cycle(board *[][]rune, southStart, westStart *int) {
	tiltNorthSouth(board, -1, 0)
	tiltEastWest(board, -1, 0)
	tiltNorthSouth(board, 1, *southStart)
	tiltEastWest(board, 1, *westStart)
}

func countRows(board *[][]rune) int {
	round := 'O'
	sum := 0
	for i, line := range *board {
		adder := len(*board) - i
		for _, r := range line {
			if r == round {
				sum += adder
			}
		}
	}
	return sum
}

func tiltNorthSouth(board *[][]rune, dir, start int) {
	round := 'O'
	ground := '.'
	height := len(*board)
	for i := start; i >= 0 && i < height; i -= dir {
		// fmt.Println("\nMoving row ", i)
		for j := range (*board)[i] {
			// for each O, check if it can roll in the direction
			iNew := i + dir
			if (*board)[i][j] == round && iNew >= 0 && iNew < height && (*board)[iNew][j] == ground {
				// find how far in that direction it can go
				for iNew+dir >= 0 && iNew+dir < height && (*board)[iNew+dir][j] == ground {
					iNew += dir
				}
				// swap
				(*board)[iNew][j] = round
				(*board)[i][j] = ground
			}
		}
		// printBoard(*board)
	}
}

func tiltEastWest(board *[][]rune, dir, start int) {
	round := 'O'
	ground := '.'
	width := len((*board)[0])
	for j := start; j >= 0 && j < width; j -= dir {
		// fmt.Println("\nMoving col ", j)
		for i := range *board {
			// for each O, check if it can roll in the direction
			jNew := j + dir
			if (*board)[i][j] == round && jNew >= 0 && jNew < width && (*board)[i][jNew] == ground {
				// find how far in that direction it can go
				for jNew+dir >= 0 && jNew+dir < width && (*board)[i][jNew+dir] == ground {
					jNew += dir
				}
				// swap
				(*board)[i][jNew] = round
				(*board)[i][j] = ground
			}
		}
		// printBoard(*board)
	}
}

func printBoard(board [][]rune) {
	for _, line := range board {
		for _, r := range line {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}

func parseInput(input string) [][]rune {
	lines := strings.Split(input, "\n")
	board := make([][]rune, len(lines))
	for i, line := range lines {
		board[i] = make([]rune, len(line))
		for j, c := range line {
			board[i][j] = c
		}
	}
	return board
}

var input2 = `fixme`

var input1 = `fixme`
