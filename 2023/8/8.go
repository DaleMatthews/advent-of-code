package main

import (
	"fmt"
	"strings"
)

var input1 = `fixme`

var input2 = `fixme`

var input3 = `fixme`

var input4 = `fixme`

type Node struct {
	str      string
	leftStr  string
	rightStr string
	left     *Node
	right    *Node
}

func main() {
	input := input1
	directions, nodes := parseInput(input)
	steps := solveForNode(directions, nodes, nodes["AAA"], true)
	fmt.Println("part1", steps)

	// directions, nodes = parseInput(input4)
	steps = part2(directions, nodes)
	fmt.Println("part2", steps)
}

func parseInput(input string) (string, map[string]*Node) {
	parts := strings.Split(input, "\n\n")
	directions := parts[0]
	lines := strings.Split(parts[1], "\n")

	nodes := make(map[string]*Node)
	for _, line := range lines {
		cur := line[:3]
		left := line[7:10]
		right := line[12:15]
		node := Node{cur, left, right, nil, nil}
		nodes[cur] = &node
	}

	for _, node := range nodes {
		node.left = nodes[node.leftStr]
		node.right = nodes[node.rightStr]
	}

	return directions, nodes
}

func solveForNode(directions string, nodes map[string]*Node, start *Node, fullZ bool) int {
	cur := start
	steps := 0
	i := 0
	for {
		direction := directions[i]
		steps++
		i++
		if i == len(directions) {
			i = 0
		}
		cur = step(direction, cur)
		if fullZ && cur.str == "ZZZ" || !fullZ && cur.str[2] == 'Z' {
			break
		}
	}
	return steps
}

func part2(directions string, nodes map[string]*Node) int {
	// find all starting nodes
	curs := make([]*Node, 0, 1)
	for _, node := range nodes {
		if node.str[2] == 'A' {
			curs = append(curs, node)
		}
	}

	solutions := make([]int, len(curs))
	for i, cur := range curs {
		solutions[i] = solveForNode(directions, nodes, cur, false)
	}
	return LCM(solutions[0], solutions[1], solutions[2:]...)
}

func step(direction byte, cur *Node) *Node {
	if direction == 'R' {
		return cur.right
	}
	return cur.left
}

// the below two funcs are from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
