package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Lens struct {
	length int
	label  string
	box    *Box
	next   *Lens
	prev   *Lens
}

type Box struct {
	start *Lens
}

var boxes = make([]Box, 256)
var allLenses = make(map[string]Lens)
var remove = '-'
var add = '='

func main() {
	input := input1
	lines := parseInput(input)
	fmt.Println("Part 1 hash result: ", calculateSum(lines))
	part2(lines)
}

func part2(lines []string) {
	for _, line := range lines {
		label, hash, op, lensLength := readLine(line)
		// fmt.Println("label: ", label, "\tbox: ", hash, "\top: ", string(op), "\tfocal: ", lensLength)
		if op == add {
			addLensToBox(label, hash, lensLength)
		} else {
			removeLensFromBox(label, hash)
		}
		// printBox(0)
		// printBox(3)
	}
	fmt.Println("Part 2 sum: ", calculateSumPart2())
}

func printBox(hash int) {
	box := &boxes[hash]
	fmt.Print("Box ", hash, ": ")
	lens := box.start
	for lens != nil {
		fmt.Print("[", lens.label, " ", lens.length, "] ")
		lens = lens.next
	}
	fmt.Println()
}

func calculateSumPart2() int {
	// The focusing power of a single lens is the result of multiplying together:
	//   - One plus the box number of the lens in question.
	//   - The slot number of the lens within the box: 1 for the first lens, 2 for the second lens, and so on.
	//   - The focal length of the lens.
	sum := 0
	for i, box := range boxes {
		lens := box.start
		j := 0
		for lens != nil {
			sum += (i + 1) * (j + 1) * lens.length
			lens = lens.next
			j++
		}
	}
	return sum
}

func addLensToBox(label string, hash int, lensLength int) {
	box := &boxes[hash]
	lens, exists := getLensFromBox(label, box)
	if exists {
		lens.length = lensLength
	} else {
		lens = &Lens{lensLength, label, box, nil, nil}
		if box.start == nil {
			box.start = lens
		} else {
			lastLens := box.start
			for lastLens.next != nil {
				lastLens = lastLens.next
			}
			lastLens.next = lens
			lens.prev = lastLens
		}
	}
}

func removeLensFromBox(label string, hash int) {
	box := &boxes[hash]
	if box.start == nil {
		return
	}
	lens, exists := getLensFromBox(label, box)
	if !exists {
		return
	}
	if box.start.label == label {
		box.start = box.start.next
		if box.start != nil {
			box.start.prev = nil
		}
	}
	prev := lens.prev
	next := lens.next
	if prev != nil && next != nil {
		prev.next = next
		next.prev = prev
	} else if prev != nil {
		prev.next = nil
	} else if next != nil {
		next.prev = nil
	}
}

func getLensFromBox(label string, box *Box) (*Lens, bool) {
	lens := box.start
	for lens != nil {
		if lens.label == label {
			return lens, true
		}
		lens = lens.next
	}
	return nil, false
}

func readLine(line string) (string, int, rune, int) {
	addIndex := strings.IndexRune(line, add)
	if addIndex != -1 {
		label := line[:addIndex]
		lensLength, _ := strconv.Atoi(line[addIndex+1:])
		return label, calculateHash(label), add, lensLength
	}
	removeIndex := strings.IndexRune(line, remove)
	label := line[:removeIndex]
	return label, calculateHash(label), remove, -1
}

func calculateSum(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += calculateHash(line)
	}
	return sum
}

func calculateHash(line string) int {
	hash := 0
	for i := range line {
		hash = ((hash + int(line[i])) * 17) % 256
	}
	return hash
}

func parseInput(input string) []string {
	return strings.Split(input, ",")
}

var input2 = `fixme`
