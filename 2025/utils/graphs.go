package utils

import (
	"math"
)

type Coord struct {
	x, y int
}

func NewCoord(x, y int) Coord {
	return Coord{x, y}
}

type Coord3D struct {
	X, Y, Z int
}

func (c1 *Coord3D) Dist(c2 Coord3D) float64 {
	dx := float64(c1.X - c2.X)
	dy := float64(c1.Y - c2.Y)
	dz := float64(c1.Z - c2.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
