package main

import (
	"fmt"
	"slices"
	"strings"
)

type Tile struct {
	char     rune
	engNorth bool
	engEast  bool
	engSouth bool
	engWest  bool
}

func main() {
	input := input1
	board := parseInput(input)
	height := len(board)
	width := len(board[0])

	sums := make([]int, 0, 2*width+2*height)
	for i := range board {
		board1 := makeBoardCopy(&board)
		traceLight(board1, i, 0, 1, 0)
		sums = append(sums, countEnergized(board1))
		board2 := makeBoardCopy(&board)
		traceLight(board2, i, width-1, -1, 0)
		sums = append(sums, countEnergized(board2))
	}
	for j := range board[0] {
		board1 := makeBoardCopy(&board)
		traceLight(board1, 0, j, 0, 1)
		sums = append(sums, countEnergized(board1))
		board2 := makeBoardCopy(&board)
		traceLight(board2, height-1, j, 0, -1)
		sums = append(sums, countEnergized(board2))
	}
	// printBoard(&board)
	fmt.Println("Part 1 num energized: ", sums[0])
	fmt.Println("Part 2 most energized: ", slices.Max(sums))
}

func makeBoardCopy(board *[][]Tile) *[][]Tile {
	boardCopy := make([][]Tile, len(*board))
	for i, tiles := range *board {
		boardCopy[i] = make([]Tile, len(tiles))
		copy(boardCopy[i], tiles)
	}
	return &boardCopy
}

func parseInput(input string) [][]Tile {
	lines := strings.Split(input, "\n")
	board := make([][]Tile, len(lines))
	for i, line := range lines {
		board[i] = make([]Tile, len(line))
		for j, r := range line {
			board[i][j] = Tile{char: r}
		}
	}
	return board
}

func countEnergized(board *[][]Tile) int {
	sum := 0
	for _, tiles := range *board {
		for _, tile := range tiles {
			if tile.engNorth || tile.engEast || tile.engSouth || tile.engWest {
				sum++
			}
		}
	}
	return sum
}

// Empty:    .
// Mirror:   / or \
// Splitter: | or -

func traceLight(board *[][]Tile, i, j, dirX, dirY int) {
	height := len(*board)
	width := len((*board)[0])
	for i >= 0 && i < height && j >= 0 && j < width {
		tile := &(*board)[i][j]
		if existingPath(tile, dirX, dirY) {
			break
		}
		energizeTile(tile, dirX, dirY)
		if tile.char == '.' || (dirX != 0 && tile.char == '-') || (dirY != 0 && tile.char == '|') {
			// nothing happens, move forward
			i = i + dirY
			j = j + dirX
		} else if tile.char == '/' || tile.char == '\\' {
			i, j, dirX, dirY = resolveMirror(tile.char, i, j, dirX, dirY)
		} else if tile.char == '|' {
			traceLight(board, i-1, j, 0, -1)
			traceLight(board, i+1, j, 0, 1)
		} else if tile.char == '-' {
			traceLight(board, i, j+1, 1, 0)
			traceLight(board, i, j-1, -1, 0)
		} else {
			panic("we missed something")
		}
	}
}

func resolveMirror(mirror rune, i, j, dirX, dirY int) (int, int, int, int) {
	if mirror == '/' {
		if dirX == 1 {
			return i - 1, j, 0, -1
		} else if dirX == -1 {
			return i + 1, j, 0, 1
		} else if dirY == 1 {
			return i, j - 1, -1, 0
		}
		return i, j + 1, 1, 0
	}
	// mirror is \
	if dirX == 1 {
		return i + 1, j, 0, 1
	} else if dirX == -1 {
		return i - 1, j, 0, -1
	} else if dirY == 1 {
		return i, j + 1, 1, 0
	}
	return i, j - 1, -1, 0
}

func existingPath(tile *Tile, dirX, dirY int) bool {
	if dirY == -1 {
		return tile.engNorth
	}
	if dirX == 1 {
		return tile.engEast
	}
	if dirY == 1 {
		return tile.engSouth
	}
	return tile.engWest
}

func energizeTile(tile *Tile, dirX, dirY int) {
	if dirY == -1 {
		tile.engNorth = true
	}
	if dirX == 1 {
		tile.engEast = true
	}
	if dirY == 1 {
		tile.engSouth = true
	}
	if dirX == -1 {
		tile.engWest = true
	}
}

func printBoard(board *[][]Tile) {
	for _, tiles := range *board {
		for _, tile := range tiles {
			energizedSum := 0
			if tile.engNorth {
				energizedSum++
			}
			if tile.engEast {
				energizedSum++
			}
			if tile.engSouth {
				energizedSum++
			}
			if tile.engWest {
				energizedSum++
			}
			if energizedSum > 1 {
				fmt.Print(energizedSum)
			} else if energizedSum == 1 {
				if tile.engNorth {
					fmt.Print("^")
				} else if tile.engEast {
					fmt.Print(">")
				} else if tile.engSouth {
					fmt.Print("v")
				} else if tile.engWest {
					fmt.Print("<")
				}
			} else {
				fmt.Print(string(tile.char))
			}
		}
		fmt.Println()
	}
}

var input2 = `fixme`

var input1 = `fixme`
