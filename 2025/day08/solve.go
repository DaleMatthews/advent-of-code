package day08

import (
	"cmp"
	"slices"
	"spissable/advent-of-go-template/utils"
	"strconv"
	"strings"
)

type Edge struct {
	c1, c2 int // index of coords array
	dist   float64
}

func SolvePuzzle1(input string, maxConnectionChecks int) int {
	coords := parseInput(input)
	edges := getEdges(coords)
	dsu := utils.NewDSU(len(coords))
	for i := 0; i < maxConnectionChecks; i++ {
		e := &edges[i]
		dsu.Union(e.c1, e.c2)
	}
	sizes := dsu.RootSizes()
	slices.SortFunc(sizes, func(a, b int) int {
		return -cmp.Compare(a, b)
	})
	return sizes[0] * sizes[1] * sizes[2]
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
			edges = append(edges, Edge{i, j, dist})
		}
	}
	slices.SortFunc(edges, func(a, b Edge) int {
		return cmp.Compare(a.dist, b.dist)
	})
	return edges
}
