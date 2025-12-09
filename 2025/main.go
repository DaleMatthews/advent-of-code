package main

import (
	"fmt"
	"os"
	"time"

	"spissable/advent-of-go-template/day09"
)

func main() {
	// Read input file
	data, err := os.ReadFile("day09/input1.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	input := string(data)

	// Run and time SolvePuzzle2
	start := time.Now()
	result := day09.SolvePuzzle2(input)
	elapsed := time.Since(start)

	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Time: %v\n", elapsed)
}
