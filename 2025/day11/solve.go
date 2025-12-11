package day11

import (
	"strings"
)

type Device struct {
	key     string
	outputs []*Device
}

func SolvePuzzle1(input string) int {
	_, m := parseInput(input)
	return fold(
		m["you"],
		func(d *Device) int { return 1 },
		func(_ *Device, counts []int) int { return sum(counts) },
	)
}

func SolvePuzzle2(input string) int {
	_, m := parseInput(input)
	return findPaths(m["svr"], false, false, make(map[cacheKey]int))
}

type cacheKey struct {
	node    *Device
	seenFFT bool
	seenDAC bool
}

func findPaths(d *Device, seenFFT, seenDAC bool, cache map[cacheKey]int) int {
	seenFFT = seenFFT || d.key == "fft"
	seenDAC = seenDAC || d.key == "dac"

	key := cacheKey{d, seenFFT, seenDAC}
	if v, ok := cache[key]; ok {
		return v
	}

	if len(d.outputs) == 0 {
		if seenFFT && seenDAC {
			return 1
		}
		return 0
	}

	count := 0
	for _, child := range d.outputs {
		count += findPaths(child, seenFFT, seenDAC, cache)
	}

	cache[key] = count
	return count
}

func parseInput(input string) ([]*Device, map[string]*Device) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	devices := make([]*Device, len(lines))
	deviceMap := make(map[string]*Device)

	for i, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), ": ")
		key := parts[0]
		device := &Device{key, []*Device{}}
		deviceMap[key] = device
		devices[i] = device
	}
	for i, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), ": ")
		rhs := strings.Split(parts[1], " ")
		for _, key := range rhs {
			if device, exists := deviceMap[key]; exists {
				devices[i].outputs = append(devices[i].outputs, device)
			}
		}
	}
	return devices, deviceMap
}

func fold[T any](node *Device, do func(*Device) T, reduce func(*Device, []T) T) T {
	if len(node.outputs) == 0 {
		return do(node)
	}

	childResults := make([]T, len(node.outputs))
	for i, child := range node.outputs {
		childResults[i] = fold(child, do, reduce)
	}
	return reduce(node, childResults)
}

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
