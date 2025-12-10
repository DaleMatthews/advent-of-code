package day10

import (
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
		numPresses := AStarPart2(make([]int, len(voltages[i])), voltages[i], machineButtons[i])
		println(numPresses)
		sum += numPresses
	}
	return sum
}

type State struct {
	voltages     []int
	validButtons [][]int
	depth        int
	h            int
}

func AStarPart2(initial []int, target []int, machineButtons [][]int) int {

	initialH := heuristic(&initial, &target)
	queue := []State{{
		voltages:     initial,
		validButtons: machineButtons,
		depth:        0,
		h:            initialH,
	}}

	visited := make(map[string]bool)
	visited[toKey(&initial)] = true

	for len(queue) > 0 {
		// Find state with lowest f = depth + h
		bestIdx := 0
		bestF := queue[0].depth + queue[0].h
		for i := 1; i < len(queue); i++ {
			f := queue[i].depth + queue[i].h
			if f < bestF {
				bestF = f
				bestIdx = i
			}
		}

		current := queue[bestIdx]
		queue = append(queue[:bestIdx], queue[bestIdx+1:]...)

		if isEqual(current.voltages, target) {
			return current.depth
		}

		for _, button := range current.validButtons {
			newVoltages := increaseVoltages(current.voltages, button)
			key := toKey(&newVoltages)

			if !visited[key] {
				visited[key] = true
				validButtons := filterValidButtons(&newVoltages, &target, &current.validButtons)
				queue = append(queue, State{
					voltages:     newVoltages,
					validButtons: validButtons,
					depth:        current.depth + 1,
					h:            heuristic(&newVoltages, &target),
				})
			}
		}
	}
	return -1
}

func heuristic(voltages, target *[]int) int {
	return sum(target) - sum(voltages)
}

func filterValidButtons(voltages, target *[]int, buttons *[][]int) [][]int {
	newButtons := make([][]int, 0, len(*buttons))
	for _, b := range *buttons {
		valid := true
		for _, i := range b {
			if (*voltages)[i] >= (*target)[i] {
				valid = false
				break
			}
		}
		if valid {
			newButtons = append(newButtons, b)
		}
	}
	return newButtons
}

func isEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func toKey(state *[]int) string {
	b := strings.Builder{}
	for i, v := range *state {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(strconv.Itoa(v))
	}
	return b.String()
}

func increaseVoltages(state []int, button []int) []int {
	newState := make([]int, len(state))
	copy(newState, state)
	for _, pos := range button {
		newState[pos]++
	}
	return newState
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

func sum(arr *[]int) int {
	sum := 0
	for _, elem := range *arr {
		sum += elem
	}
	return sum
}
