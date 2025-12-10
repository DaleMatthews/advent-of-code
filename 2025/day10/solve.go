package day10

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) int {
	numLights, schematics, machineButtons := parseInput1(input)
	sum := 0
	for i := 0; i < len(schematics); i++ {
		numPresses := BFSPart1(0, schematics[i], machineButtons[i])
		println("[", machineToString(schematics[i], numLights[i]), "]", numPresses)
		sum += numPresses
	}
	return sum
}

func BFSPart1(initial int, target int, machineButtons []int) int {
	queue := []int{initial}
	visited := make(map[int]bool)
	visited[initial] = true
	depth := 0
	for len(queue) > 0 {
		statesToCheck := len(queue)
		for i := 0; i < statesToCheck; i++ {
			state := queue[0]
			queue = queue[1:]
			if state == target {
				return depth
			}

			for _, button := range machineButtons {
				newState := toggleLights(state, button)
				if !visited[newState] {
					visited[newState] = true
					queue = append(queue, newState)
				}
			}
		}
		depth++
	}
	return -1
}

func toggleLights(lights int, button int) int {
	return lights ^ button
}

// uses bit masks for machine states
func parseInput1(input string) ([]int, []int, [][]int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	numLights := make([]int, len(lines))
	schematics := make([]int, len(lines))
	buttons := make([][]int, len(lines))

	for i, line := range lines {
		line = strings.TrimSpace(line)
		// schematic
		states := line[1:strings.Index(line, "]")]
		numLights[i] = len(states)
		for j, s := range states {
			if s == '#' {
				schematics[i] |= 1 << (len(states) - 1 - j) // set bit
				// schematics[i] &^= 1 << i // clear bit
			}
		}

		// buttons
		allButtons := strings.Split(line[strings.Index(line, "]")+3:strings.Index(line, ") {")], ") (")
		buttons[i] = make([]int, len(allButtons))
		for j, buttonsStr := range allButtons {
			nums := strings.Split(buttonsStr, ",")
			for _, num := range nums {
				parsed, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				buttons[i][j] |= 1 << (numLights[i] - 1 - parsed)
			}
		}
	}
	return numLights, schematics, buttons
}

func machineToString(lights, length int) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		if lights&(1<<(length-1-i)) != 0 {
			b[i] = '#'
		} else {
			b[i] = '.'
		}
	}
	return string(b)
}

func SolvePuzzle2(input string) int {
	machineButtons, voltages := parseInput2(input)
	sum := 0
	for i := 0; i < len(voltages); i++ {
		numPresses := getMinPresses2(voltages[i], machineButtons[i])
		println(numPresses)
		sum += numPresses
	}
	return sum
}

func getMinPresses2(target []int, buttons [][]int) int {
	numButtons := len(buttons)
	numIndices := len(target)

	// Build coefficient matrix
	coeff := make([][]int, numIndices)
	for i := range coeff {
		coeff[i] = make([]int, numButtons)
	}
	for j, button := range buttons {
		for _, idx := range button {
			coeff[idx][j]++
		}
	}

	// Write GMPL/MathProg file
	var sb strings.Builder
	for j := 0; j < numButtons; j++ {
		fmt.Fprintf(&sb, "var x%d, integer, >= 0;\n", j)
	}
	sb.WriteString("\n")

	sb.WriteString("minimize total: ")
	for j := 0; j < numButtons; j++ {
		if j > 0 {
			sb.WriteString(" + ")
		}
		fmt.Fprintf(&sb, "x%d", j)
	}
	sb.WriteString(";\n\n")

	for i := 0; i < numIndices; i++ {
		fmt.Fprintf(&sb, "s.t. c%d: ", i)
		first := true
		for j := 0; j < numButtons; j++ {
			if coeff[i][j] > 0 {
				if !first {
					sb.WriteString(" + ")
				}
				if coeff[i][j] == 1 {
					fmt.Fprintf(&sb, "x%d", j)
				} else {
					fmt.Fprintf(&sb, "%d*x%d", coeff[i][j], j)
				}
				first = true
				first = false
			}
		}
		fmt.Fprintf(&sb, " = %d;\n", target[i])
	}

	sb.WriteString("\nsolve;\n\n")
	sb.WriteString("printf \"RESULT: %d\\n\", ")
	for j := 0; j < numButtons; j++ {
		if j > 0 {
			sb.WriteString(" + ")
		}
		fmt.Fprintf(&sb, "x%d", j)
	}
	sb.WriteString(";\n\nend;\n")

	// Write to temp file
	f, _ := os.CreateTemp("", "problem*.mod")
	f.WriteString(sb.String())
	f.Close()

	// Run GLPK
	out, err := exec.Command("glpsol.exe", "--math", f.Name()).Output()
	os.Remove(f.Name())

	if err != nil {
		return -1
	}

	// Parse result
	for _, line := range strings.Split(string(out), "\n") {
		if strings.HasPrefix(line, "RESULT: ") {
			val, _ := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "RESULT: ")))
			return val
		}
	}

	return -1
}

func parseInput2(input string) ([][][]int, [][]int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	buttons := make([][][]int, len(lines))
	voltages := make([][]int, len(lines))

	for i, line := range lines {
		line = strings.TrimSpace(line)

		// buttons
		allButtons := strings.Split(line[strings.Index(line, "]")+3:strings.Index(line, ") {")], ") (")
		buttons[i] = make([][]int, len(allButtons))
		for j, buttonsStr := range allButtons {
			nums := strings.Split(buttonsStr, ",")
			buttons[i][j] = make([]int, len(nums))
			for k, num := range nums {
				parsed, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				buttons[i][j][k] = parsed
			}
		}

		// voltages
		nums := strings.Split(line[strings.Index(line, "{")+1:len(line)-1], ",")
		voltages[i] = make([]int, len(nums))
		for j, num := range nums {
			parsed, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			voltages[i][j] = parsed
		}
	}
	return buttons, voltages
}
