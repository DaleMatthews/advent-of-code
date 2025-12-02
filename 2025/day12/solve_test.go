package day12_test

import (
	"spissable/advent-of-go-template/day12"
	"spissable/advent-of-go-template/utils"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t, "input1.txt")
	result := day12.SolvePuzzle1(input)
	utils.LogResult(t, result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t, "input2.txt")
	result := day12.SolvePuzzle2(input)
	utils.LogResult(t, result)
}
