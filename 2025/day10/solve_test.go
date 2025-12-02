package day10_test

import (
	"spissable/advent-of-go-template/day10"
	"spissable/advent-of-go-template/utils"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t, "input1.txt")
	result := day10.SolvePuzzle1(input)
	utils.LogResult(t, result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t, "input2.txt")
	result := day10.SolvePuzzle2(input)
	utils.LogResult(t, result)
}
