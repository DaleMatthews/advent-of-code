package day01_test

import (
	"spissable/advent-of-go-template/day01"
	"spissable/advent-of-go-template/utils"
	"testing"
)

func TestSolveExample1(t *testing.T) {
	input := utils.ReadInput(t, "example1.txt")
	result := day01.SolvePuzzle1(input)
	utils.LogResult(t, result)
}

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t, "input1.txt")
	result := day01.SolvePuzzle1(input)
	utils.LogResult(t, result)
}

func TestSolveExample2(t *testing.T) {
	input := utils.ReadInput(t, "example2.txt")
	result := day01.SolvePuzzle1(input)
	utils.LogResult(t, result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t, "input2.txt")
	result := day01.SolvePuzzle2(input)
	utils.LogResult(t, result)
}
