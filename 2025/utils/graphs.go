package utils

import (
	"math"
	"strconv"
)

type Coord struct {
	x, y int
}

func NewCoord(x, y int) Coord {
	return Coord{x, y}
}

type Coord3D struct {
	x, y, z int
}

func New3DCoord(x, y, z int) Coord3D {
	return Coord3D{x, y, z}
}

func Dist3D(c1 Coord3D, c2 Coord3D) float64 {
	return math.Sqrt(math.Pow(float64(c1.x)-float64(c2.x), 2) + math.Pow(float64(c1.y)-float64(c2.y), 2) + math.Pow(float64(c1.z)-float64(c2.z), 2))
}

func Coord3DToString(c Coord3D) string {
	return strconv.Itoa(c.x) + "," + strconv.Itoa(c.y) + "," + strconv.Itoa(c.z)
}
