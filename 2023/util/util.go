package util

import (
	"fmt"
	"strconv"
)

func StringsToFloats(strs []string) []float64 {
	float64s := make([]float64, len(strs))
	for i, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		float64s[i] = float64(num)
	}
	return float64s
}

func StringsToInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints[i] = num
	}
	return ints
}

func PrintStringArray(strs []string) {
	for _, str := range strs {
		fmt.Println(str)
	}
}

func PrintBoard[T any](board [][]T, separator string) {
	for _, line := range board {
		for _, elem := range line {
			fmt.Print(elem, separator)
		}
		fmt.Println()
	}
}

func Sum(arr []int) int {
	sum := 0
	for _, elem := range arr {
		sum += elem
	}
	return sum
}
