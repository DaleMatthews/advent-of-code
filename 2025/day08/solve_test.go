package day08_test

import (
	"spissable/advent-of-go-template/day08"
	"spissable/advent-of-go-template/utils"
	"testing"
)

func TestSolveExample1(t *testing.T) {
	input := utils.ReadInput(t, "example1.txt")
	result := day08.SolvePuzzle1(input)
	utils.LogResult(t, result)
}

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t, "input1.txt")
	result := day08.SolvePuzzle1(input)
	utils.LogResult(t, result)
}

func TestSolveExample2(t *testing.T) {
	input := utils.ReadInput(t, "example1.txt")
	result := day08.SolvePuzzle2(input)
	utils.LogResult(t, result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t, "input1.txt")
	result := day08.SolvePuzzle2(input)
	utils.LogResult(t, result)
}

func BenchmarkSolvePuzzle2(b *testing.B) {
	input := utils.ReadInput(nil, "input1.txt") // might need adjustment
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day08.SolvePuzzle2(input)
	}
}
