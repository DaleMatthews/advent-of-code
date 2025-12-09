package day09

import (
	"spissable/advent-of-go-template/utils"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) int {
	coords := parseInput(input)
	max := -1
	var maxC1, maxC2 utils.Coord2D
	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			area := abs(coords[i].X-coords[j].X) * abs(coords[i].Y-coords[j].Y)
			if area > max {
				max = area
				maxC1 = coords[i]
				maxC2 = coords[j]
			}
		}
	}
	println("c1:", maxC1.X, ",", maxC1.Y)
	println("c2:", maxC2.X, ",", maxC2.Y)
	return (abs(maxC1.X-maxC2.X) + 1) * (abs(maxC1.Y-maxC2.Y) + 1)
}

func SolvePuzzle2(input string) int {
	coords := parseInput(input)
	allEdgeCoords := make(map[utils.Coord2D]bool)
	// build cache of all edge coords
	for i := 0; i < len(coords); i++ {
		j := i + 1
		if j == len(coords) {
			j = 0
		}
		allEdgeCoords[coords[i]] = true
		allEdgeCoords[coords[j]] = true
		for c := range utils.CoordsInStraightLine(coords[i], coords[j]) {
			allEdgeCoords[c] = true
		}
	}
	// find max rectangle where the boundaries never fully cross a polygon edge
	max := -1
	var maxC1, maxC2 utils.Coord2D
	for i := 0; i < len(coords)-1; i++ {
		println("Checking line", i+1)
		for j := i + 1; j < len(coords); j++ {
			if coords[i].X == coords[j].X || coords[i].Y == coords[j].Y {
				continue // let's just ignore the 1-width rectangles for simplicity...
			}
			if !rectIsInPoly(coords[i], coords[j], &allEdgeCoords) {
				continue
			}
			area := (abs(maxC1.X-maxC2.X) + 1) * (abs(maxC1.Y-maxC2.Y) + 1)
			if area > max {
				max = area
				maxC1 = coords[i]
				maxC2 = coords[j]
				println("c1:", maxC1.X, ",", maxC1.Y)
				println("c2:", maxC2.X, ",", maxC2.Y)
				println("area:", area)
			}
		}
	}
	println("c1:", maxC1.X, ",", maxC1.Y)
	println("c2:", maxC2.X, ",", maxC2.Y)
	area := (abs(maxC1.X-maxC2.X) + 1) * (abs(maxC1.Y-maxC2.Y) + 1)
	println("area:", area)
	return area
}

func rectIsInPoly(c1, c2 utils.Coord2D, edges *map[utils.Coord2D]bool) bool {
	// c1----line1----*
	// |              |
	// |line4    line2|
	// |              |
	// *----line3----c2

	var c utils.Coord2D
	needsCheck := true
	// line 1
	for c = range utils.StepTowardsX(c1, c2) {
		_, isEdge := (*edges)[c]
		if isEdge {
			needsCheck = true
		} else {
			if needsCheck && !coordIsInPoly(c, edges) {
				return false
			}
			needsCheck = false
		}
	}
	// line 2
	for c = range utils.StepTowardsY(c, c2) {
		_, isEdge := (*edges)[c]
		if isEdge {
			needsCheck = true
		} else {
			if needsCheck && !coordIsInPoly(c, edges) {
				return false
			}
			needsCheck = false
		}
	}
	// line 3
	for c = range utils.StepTowardsX(c2, c1) {
		_, isEdge := (*edges)[c]
		if isEdge {
			needsCheck = true
		} else {
			if needsCheck && !coordIsInPoly(c, edges) {
				return false
			}
			needsCheck = false
		}
	}
	// line 4
	for c = range utils.StepTowardsY(c, c1) {
		_, isEdge := (*edges)[c]
		if isEdge {
			needsCheck = true
		} else {
			if needsCheck && !coordIsInPoly(c, edges) {
				return false
			}
			needsCheck = false
		}
	}
	return true
}

var cache = make(map[utils.Coord2D]bool)

func coordIsInPoly(coord utils.Coord2D, edges *map[utils.Coord2D]bool) bool {
	val, exists := cache[coord]
	if exists {
		return val
	}
	numEdges := 0
	inEdge := false
	if coord.X > 80000 {
		// cast right
		c := utils.Coord2D{X: coord.X + 1, Y: coord.Y}
		for c.X < 98380 {
			if _, exists := (*edges)[c]; exists {
				inEdge = true
			} else if inEdge {
				numEdges++
				inEdge = false
			}
			c = utils.Coord2D{X: c.X + 1, Y: c.Y}
		}
	} else if coord.Y > 80000 {
		// cast down
		c := utils.Coord2D{X: coord.X, Y: coord.Y + 1}
		for c.Y < 98380 {
			if _, exists := (*edges)[c]; exists {
				inEdge = true
			} else if inEdge {
				numEdges++
				inEdge = false
			}
			c = utils.Coord2D{X: c.X, Y: c.Y + 1}
		}
	} else if coord.Y < 20000 {
		// cast up
		c := utils.Coord2D{X: coord.X, Y: coord.Y - 1}
		for c.Y >= 0 {
			if _, exists := (*edges)[c]; exists {
				inEdge = true
			} else if inEdge {
				numEdges++
				inEdge = false
			}
			c = utils.Coord2D{X: c.X, Y: c.Y - 1}
		}
	} else {
		// cast left
		c := utils.Coord2D{X: coord.X - 1, Y: coord.Y}
		for c.X >= 0 {
			if _, exists := (*edges)[c]; exists {
				inEdge = true
			} else if inEdge {
				numEdges++
				inEdge = false
			}
			c = utils.Coord2D{X: c.X - 1, Y: c.Y}
		}
	}
	isInPoly := numEdges%2 == 1
	cache[coord] = isInPoly
	return isInPoly
}

func parseInput(input string) []utils.Coord2D {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	coords := make([]utils.Coord2D, len(lines))

	for i, line := range lines {
		numStrs := strings.Split(strings.TrimSpace(line), ",")
		lhs, _ := strconv.Atoi(numStrs[0])
		rhs, _ := strconv.Atoi(numStrs[1])
		coords[i] = utils.Coord2D{X: lhs, Y: rhs}
	}
	return coords
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
