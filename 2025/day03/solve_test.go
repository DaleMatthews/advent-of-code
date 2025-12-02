package day03_test

import (
	"spissable/advent-of-go-template/day03"
	"spissable/advent-of-go-template/utils"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t, "input1.txt")
	result := day03.SolvePuzzle1(input)
	utils.LogResult(t, result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t, "input1.txt")
	result := day03.SolvePuzzle2(input)
	utils.LogResult(t, result)
}

func BenchmarkSolvePuzzle2(b *testing.B) {
	input := utils.ReadInput(nil, "input1.txt") // might need adjustment
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day03.SolvePuzzle2(input)
	}
}
