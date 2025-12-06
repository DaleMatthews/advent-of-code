package day06

import (
	"spissable/advent-of-go-template/utils"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func solve(problems [][]int, ops []string) int {
	total := 0
	for i, problem := range problems {
		var sol int
		if ops[i] == "+" {
			sol = lo.Sum(problem)
		} else {
			sol = lo.Product(problem)
		}
		// println("sol:", sol)
		total += sol
	}

	return total
}

func SolvePuzzle1(input string) int {
	problems, ops := parseInput1(input)
	return solve(problems, ops)
}

func SolvePuzzle2(input string) int {
	problems, ops := parseInput2(input)
	return solve(problems, ops)
}

func parseInput1(input string) ([][]int, []string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	board := make([][]int, len(lines)-1)

	for i, line := range lines[:len(lines)-1] {
		fields := strings.Fields(line)
		board[i] = make([]int, len(fields))
		for j, str := range fields {
			board[i][j], _ = strconv.Atoi(str)
		}
	}
	return utils.Transpose(board), strings.Fields(lines[len(lines)-1])
}

func parseInput2(input string) ([][]int, []string) {
	lines := strings.Split(input, "\r\n")
	lines = lines[:len(lines)-1] // remove empty line
	board := make([][]string, len(lines))
	for i, line := range lines {
		board[i] = make([]string, len(line))
		for j, r := range line {
			board[i][j] = string(r)
		}
	}

	board = utils.Transpose(board)
	problems := make([][]int, 0)
	ops := make([]string, 0)
	problem := make([]int, 0)
	for _, col := range board {
		numStr := ""
		opCol := false
		for _, s := range col {
			if s == "*" || s == "+" {
				opCol = true
				ops = append(ops, s)
			} else if s != " " {
				numStr += s
			}
		}
		if opCol {
			if len(problem) > 0 {
				problems = append(problems, problem)
			}
			problem = make([]int, 0)
		}
		if numStr == "" {
			continue
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic("cant parse line")
		}
		problem = append(problem, num)
	}
	if len(problem) > 0 {
		problems = append(problems, problem)
	}

	return problems, ops
}
