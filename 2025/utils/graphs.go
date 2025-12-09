package utils

import (
	"iter"
	"math"
)

type Coord2D struct {
	X, Y int
}

/** must be straight line */
func CoordsInStraightLine(c1 Coord2D, c2 Coord2D) iter.Seq[Coord2D] {
	return func(yield func(Coord2D) bool) {
		dx := sign(c2.X - c1.X)
		dy := sign(c2.Y - c1.Y)
		c := Coord2D{c1.X + dx, c1.Y + dy}
		for c != c2 {
			if !yield(c) {
				return
			}
			c.X += dx
			c.Y += dy
		}
	}
}

func StepTowardsX(c1 Coord2D, c2 Coord2D) iter.Seq[Coord2D] {
	return func(yield func(Coord2D) bool) {
		dx := sign(c2.X - c1.X)
		c := Coord2D{c1.X + dx, c1.Y}
		for c.X != c2.X {
			if !yield(c) {
				return
			}
			c.X += dx
		}
		yield(c)
	}
}

func StepTowardsY(c1 Coord2D, c2 Coord2D) iter.Seq[Coord2D] {
	return func(yield func(Coord2D) bool) {
		dy := sign(c2.Y - c1.Y)
		c := Coord2D{c1.X, c1.Y + dy}
		for c.Y != c2.Y {
			if !yield(c) {
				return
			}
			c.Y += dy
		}
		yield(c)
	}
}

func RectCorners(c1, c2 Coord2D) (Coord2D, Coord2D) {
	return Coord2D{c1.X, c2.Y}, Coord2D{c2.X, c1.Y}
}

func sign(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
