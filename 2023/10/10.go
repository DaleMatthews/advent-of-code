package main

import (
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

var input1 string = `fixme`

var input2 string = `fixme`

var input3 string = `fixme`

var input4 string = `fixme`

var input5 string = `fixme`

var input6 string = `fixme`

var input7 string = `fixme`

type Tile struct {
	pipe     rune
	row      int
	col      int
	distance int
	forward  *Tile
	backward *Tile
}

// | - L J 7 F
func (t *Tile) ValidNorth() bool {
	return t.pipe == 'S' || t.pipe == '|' || t.pipe == '7' || t.pipe == 'F'
}

func (t *Tile) ValidEast() bool {
	return t.pipe == 'S' || t.pipe == '-' || t.pipe == '7' || t.pipe == 'J'
}

func (t *Tile) ValidSouth() bool {
	return t.pipe == 'S' || t.pipe == '|' || t.pipe == 'L' || t.pipe == 'J'
}

func (t *Tile) ValidWest() bool {
	return t.pipe == 'S' || t.pipe == '-' || t.pipe == 'L' || t.pipe == 'F'
}

func main() {
	// 393 is too high
	// 392 is wrong
	// 384 is wrong
	// 317 is correct
	input := input1
	maze, start := parseInput(input)
	// fmt.Println(startLine, startCol)
	loop := mapset.NewSet[*Tile]()
	loop.Add(start)
	end := walkMaze(maze, loop, start, nil)
	fmt.Println("Distance: ", end.distance)
	fmt.Println("Length: ", loop.Cardinality())
	replaceStartPipe(maze, start)
	inside := nodesInside(maze, loop)
	fmt.Println("Enclosed: ", inside.Cardinality())
	// printResult(maze, loop, inside)
}

func parseInput(input string) ([][]*Tile, *Tile) {
	lines := strings.Split(input, "\n")
	// util.PrintStringArray(lines)
	var start *Tile

	maze := make([][]*Tile, len(lines))
	for i, line := range lines {
		maze[i] = make([]*Tile, len(line))
		for j, c := range line {
			maze[i][j] = &Tile{
				pipe:     c,
				row:      i,
				col:      j,
				distance: -1,
				forward:  nil,
				backward: nil,
			}
			if c == 'S' {
				start = maze[i][j]
				start.distance = 0
			}
		}
	}

	return maze, start
}

func walkMaze(maze [][]*Tile, loop mapset.Set[*Tile], start1, start2 *Tile) *Tile {
	queue := getForwardNeighbors(maze, start1)
	if start2 != nil {
		queue = append(queue, getForwardNeighbors(maze, start2)...)
	}
	if len(queue) == 0 {
		return start1
	}
	for _, t := range queue {
		loop.Add(t)
	}
	queue[0].distance = start1.distance + 1
	queue[0].backward = start1
	start1.forward = queue[0]
	if len(queue) == 1 {
		return queue[0]
	}
	if start2 != nil {
		queue[1].distance = start2.distance + 1
		queue[1].backward = start2
		start2.forward = queue[1]
	} else {
		queue[1].distance = start1.distance + 1
		queue[1].backward = start1
	}
	return walkMaze(maze, loop, queue[0], queue[1])
}

func getForwardNeighbors(maze [][]*Tile, target *Tile) []*Tile {
	// maze[row][col]
	// .....
	// .S-7.
	// .|.|.
	// .L-J.
	// .....
	neighbors := make([]*Tile, 0, 2)
	var north = northTile(maze, target)
	var east = eastTile(maze, target)
	var south = southTile(maze, target)
	var west = westTile(maze, target)
	if north != nil && north.ValidNorth() && target.ValidSouth() && north.distance == -1 {
		neighbors = append(neighbors, north)
	}
	if east != nil && east.ValidEast() && target.ValidWest() && east.distance == -1 {
		neighbors = append(neighbors, east)
	}
	if south != nil && south.ValidSouth() && target.ValidNorth() && south.distance == -1 {
		neighbors = append(neighbors, south)
	}
	if west != nil && west.ValidWest() && target.ValidEast() && west.distance == -1 {
		neighbors = append(neighbors, west)
	}
	return neighbors
}

func northTile(maze [][]*Tile, target *Tile) *Tile {
	if target.row-1 != -1 {
		return maze[target.row-1][target.col]
	}
	return nil
}

func eastTile(maze [][]*Tile, target *Tile) *Tile {
	if target.col+1 != len(maze[target.row]) {
		return maze[target.row][target.col+1]
	}
	return nil
}

func southTile(maze [][]*Tile, target *Tile) *Tile {
	if target.row+1 != len(maze) {
		return maze[target.row+1][target.col]
	}
	return nil
}

func westTile(maze [][]*Tile, target *Tile) *Tile {
	if target.col-1 != -1 {
		return maze[target.row][target.col-1]
	}
	return nil
}

func replaceStartPipe(maze [][]*Tile, start *Tile) {
	// | - L J 7 F
	north := northTile(maze, start).ValidNorth()
	east := eastTile(maze, start).ValidEast()
	south := southTile(maze, start).ValidSouth()
	west := westTile(maze, start).ValidWest()
	if north && south {
		start.pipe = '|'
	} else if north && east {
		start.pipe = 'L'
	} else if north && west {
		start.pipe = 'J'
	} else if east && west {
		start.pipe = '-'
	} else if south && west {
		start.pipe = '7'
	} else {
		start.pipe = 'F'
	}
}

// part 2
// for each tile, cast a ray and count the number of times it intersects the path
// if it's even, the tile is outside
// if it's odd, the tile is inside

// ! apparently Pick's theorem would have been easier
func nodesInside(maze [][]*Tile, loop mapset.Set[*Tile]) mapset.Set[*Tile] {
	inside := mapset.NewSet[*Tile]()
	for _, row := range maze {
		for _, t := range row {
			if loop.Contains(t) {
				continue
			}
			// castRayEast worked on all example inputs but not the actual input.
			// Combining castRayEast and castRayWest to check for either, was a good enough bandaid :/
			numTouchesEast := castRayEast(maze, loop, t)
			if numTouchesEast%2 != 0 {
				// fmt.Println("(", i, ", ", j, "): inside ", numTouchesEast)
				inside.Add(t)
			}
		}
	}
	return inside
}

func castRayEast(maze [][]*Tile, loop mapset.Set[*Tile], target *Tile) int {
	count := 0
	for target != nil {
		target = eastTile(maze, target)
		if target != nil && loop.Contains(target) {
			// We found a piece of the loop.
			// If it's a | then we intersect it, otherwise we're walking along a section
			// of pipe. If we're walking a section of pipe, it's only an intersection if the start
			// and end go in opposite directions. Since 2 and 0 result in the same % 2 calculation,
			// just keep counting up for all valid north pipes. That means if ther's two that point up,
			// the pipe isn't intersected, but if one points up and the other doesn't, it becomes odd-numbered.
			// In our input, the 'S' does go up, so count it :shrug:
			if target.pipe == '|' || target.pipe == 'J' || target.pipe == 'L' {
				// this is just perpendicular portion of the loop, so it does intersect
				count++
			}
		}
	}
	return count
}

func printResult(maze [][]*Tile, loop mapset.Set[*Tile], inside mapset.Set[*Tile]) {
	for _, row := range maze {
		for _, t := range row {
			if loop.Contains(t) {
				fmt.Print("🟩")
			} else if inside.Contains(t) {
				fmt.Print("⬛")
			} else {
				fmt.Print("🟦")
			}
		}
		fmt.Println()
	}
}
