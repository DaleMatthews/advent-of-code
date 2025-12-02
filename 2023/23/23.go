package main

import (
	"1/util"
	"strings"
)

func main() {
	input := input1
	lines := parseInput(input)
	util.PrintStringArray(lines)
}

func parseInput(input string) []string {
	lines := strings.Split(input, "\n")
	return lines
}

var input1 = `fixme`
var input2 = `fixme`
