package utils

type Coord struct {
	x, y int
}

func NewCoord(x, y int) Coord {
	return Coord{x, y}
}
