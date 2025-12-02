package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Command struct {
	dir    string
	length int
	color  string
}

type Tile struct {
	color string
	dug   bool
}

var grid [][]Tile

func main() {
	input := input1 // 25157 is too low
	iOrigin, jOrigin, width, height, commands := parseInput(input)
	generateGrid(iOrigin, jOrigin, width, height, commands)
	fmt.Println("Width: ", width, "\tHeight: ", height)
	// fmt.Println(commands)
	// printGrid()
	fmt.Println("Area: ", countArea())
}

func parseInput(input string) (int, int, int, int, []Command) {
	lines := strings.Split(input, "\n")
	iMin, iMax, jMin, jMax := 0.0, 0.0, 0.0, 0.0
	commands := make([]Command, 0)
	jCur := 0.0
	iCur := 0.0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		dir := parts[0]
		length, _ := strconv.Atoi(parts[1])
		color := parts[2][1 : len(parts[2])-1]
		commands = append(commands, Command{
			dir:    dir,
			length: length,
			color:  color,
		})
		if dir == "R" {
			jCur += float64(length)
			if jCur > jMax {
				jMax = jCur
			}
		} else if dir == "L" {
			jCur -= float64(length)
			if jCur < jMin {
				jMin = jCur
			}
		} else if dir == "D" {
			iCur += float64(length)
			if iCur > iMax {
				iMax = iCur
			}
		} else if dir == "U" {
			iCur -= float64(length)
			if iCur < iMin {
				iMin = iCur
			}
		}
	}
	fmt.Println(iCur, jCur)
	width := int(jMax-jMin) + 1
	height := int(iMax-iMin) + 1
	iOrigin := int(math.Abs(iMin))
	jOrigin := int(math.Abs(jMin))
	return iOrigin, jOrigin, width, height, commands
}

func countArea() int {
	area := 0
	outside := true
	for i := 0; i < len(grid); i++ {
		seenInsideDirt := false
		lastSeen := false
		inARow := 0
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j].dug {
				inARow++
				area++
				outside = seenInsideDirt
			} else {
				inARow = 0
				if !outside {
					area++
					seenInsideDirt = true
				}
			}
		}
	}
	return area
}

func generateGrid(iOrigin, jOrigin, width, height int, commands []Command) {
	grid = make([][]Tile, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]Tile, width)
	}
	iCur := iOrigin
	jCur := jOrigin
	for _, cmd := range commands {
		iDir, jDir, iLength, jLength := 0, 0, 0, 0
		if cmd.dir == "R" {
			jDir = 1
			jLength = cmd.length
		} else if cmd.dir == "L" {
			jDir = -1
			jLength = cmd.length
		} else if cmd.dir == "D" {
			iDir = 1
			iLength = cmd.length
		} else if cmd.dir == "U" {
			iDir = -1
			iLength = cmd.length
		}
		for i, j := 0, 0; i < iLength || j < jLength; {
			grid[iCur][jCur].dug = true
			i++
			j++
			iCur += iDir
			jCur += jDir
		}
	}
}

func printGrid() {
	for _, row := range grid {
		for _, tile := range row {
			if tile.dug {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

var input2 = `fixme`

var input1 = `fixme`
