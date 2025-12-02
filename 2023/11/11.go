package main

import (
	"fmt"
	"strings"
)

var input1 string = `fixme`

var input2 string = `fixme`

var input3 string = `fixme``

type Coord struct {
	i int
	j int
}

func main() {
	input := input1
	lines, coords := parseInput(input)
	pairs := generatePairs(coords)
	fmt.Println("Num pairs: ", len(pairs))
	// part 1
	costs := calculateCosts(lines, coords, 2)
	sum := 0
	for _, pair := range pairs {
		// fmt.Println(pair[0], " --> ", pair[1])
		sum += manhattanDistance(costs, pair[0].i, pair[0].j, pair[1].i, pair[1].j)
	}
	fmt.Println("Distance sum part 1: ", sum)
	// part 2
	costs = calculateCosts(lines, coords, 1000000)
	sum = 0
	for _, pair := range pairs {
		// fmt.Println(pair[0], " --> ", pair[1])
		sum += manhattanDistance(costs, pair[0].i, pair[0].j, pair[1].i, pair[1].j)
	}
	fmt.Println("Distance sum part 2: ", sum)
}

func parseInput(input string) ([]string, []Coord) {
	lines := strings.Split(input, "\n")
	return lines, findGalaxies(lines)
}

func findGalaxies(lines []string) []Coord {
	coords := make([]Coord, 0)
	for i, line := range lines {
		for j, c := range line {
			if c == '#' {
				coords = append(coords, Coord{i, j})
			}
		}
	}
	return coords
}

func calculateCosts(lines []string, coords []Coord, emptyMultiplier int) [][]int {
	rowsTouched := make([]bool, len(lines))
	colsTouched := make([]bool, len(lines[0]))

	for _, galaxy := range coords {
		rowsTouched[galaxy.i] = true
		colsTouched[galaxy.j] = true
	}

	costs := make([][]int, len(lines))
	// fmt.Println("rows", rowsTouched)
	// fmt.Println("cols", colsTouched)

	for i, line := range lines {
		costs[i] = make([]int, len(line))
		for j := range line {
			if rowsTouched[i] && colsTouched[j] {
				costs[i][j] = 1
			} else {
				costs[i][j] = emptyMultiplier
			}
		}
	}

	return costs
}

func generatePairs(coords []Coord) [][]Coord {
	pairs := make([][]Coord, 0, len(coords))
	for i := range coords {
		for j := i + 1; j < len(coords); j++ {
			pair := make([]Coord, 2)
			pair[0] = coords[i]
			pair[1] = coords[j]
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func manhattanDistance(costs [][]int, i1, j1, i2, j2 int) int {
	distance := 0
	iStep := 1
	if i2 < i1 {
		iStep = -1
	}
	for i := i1; i != i2; i += iStep {
		distance += costs[i][j1]
	}
	jStep := 1
	if j2 < j1 {
		jStep = -1
	}
	for j := j1; j != j2; j += jStep {
		distance += costs[i2][j]
	}
	return distance
}
