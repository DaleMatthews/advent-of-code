package day08

import (
	"cmp"
	"maps"
	"slices"
	"spissable/advent-of-go-template/utils"
	"strconv"
	"strings"
)

// type Junction struct {
// 	coord           utils.Coord3D
// 	closestJunction *Junction
// 	dist            float64
// 	circuit         *map[*Junction]bool
// }

type Edge struct {
	c1, c2 *utils.Coord3D
	dist   float64
}

func SolvePuzzle1(input string, maxConnectionChecks int) int {
	coords := parseInput(input)
	edges := getEdges(coords)
	coordToCircuit := make(map[*utils.Coord3D]*map[*utils.Coord3D]bool)
	circuits := make(map[*map[*utils.Coord3D]bool]bool, 0)
	numConnections := 0
	for i := 0; i < maxConnectionChecks; i++ {
		e := &edges[i]
		if coordToCircuit[e.c1] != nil && coordToCircuit[e.c2] != nil {
			circuit1 := coordToCircuit[e.c1]
			circuit2 := coordToCircuit[e.c2]
			if circuit1 != circuit2 {
				numConnections++
				// circuit1 aborbs circuit2
				for c := range maps.Keys(*circuit2) {
					(*circuit1)[c] = true
					coordToCircuit[c] = circuit1
				}
				delete(circuits, circuit2)
			}
			continue
		}
		numConnections++
		if coordToCircuit[e.c1] != nil {
			circuit := coordToCircuit[e.c1]
			(*circuit)[e.c2] = true
			coordToCircuit[e.c2] = circuit
		} else if coordToCircuit[e.c2] != nil {
			circuit := coordToCircuit[e.c2]
			(*circuit)[e.c1] = true
			coordToCircuit[e.c1] = circuit
		} else {
			circuit := make(map[*utils.Coord3D]bool)
			circuit[e.c1] = true
			circuit[e.c2] = true
			coordToCircuit[e.c1] = &circuit
			coordToCircuit[e.c2] = &circuit
			circuits[&circuit] = true
		}
	}
	circuitList := make([]*map[*utils.Coord3D]bool, 0, len(circuits))
	for circuit := range circuits {
		circuitList = append(circuitList, circuit)
	}
	slices.SortFunc(circuitList, func(a, b *map[*utils.Coord3D]bool) int {
		return len(*b) - len(*a)
	})
	return len(*circuitList[0]) * len(*circuitList[1]) * len(*circuitList[2])
}

func SolvePuzzle2(input string) int {
	coords := parseInput(input)
	return len(coords)
}

func parseInput(input string) []utils.Coord3D {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	coords := make([]utils.Coord3D, len(lines))

	for i, line := range lines {
		nums := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		z, _ := strconv.Atoi(nums[2])
		coords[i] = utils.New3DCoord(x, y, z)
	}
	return coords
}

func getEdges(coords []utils.Coord3D) []Edge {
	edges := make([]Edge, 0)
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			dist := utils.Dist3D(coords[i], coords[j])
			edges = append(edges, Edge{&coords[i], &coords[j], dist})
		}
	}
	slices.SortFunc(edges, func(a, b Edge) int {
		return cmp.Compare(a.dist, b.dist)
	})
	return edges
}
