package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

var input1 string = `fixme`

var input2 string = `fixme`

type Range struct {
	sourceStart      int
	sourceEnd        int
	destinationStart int
	destinationEnd   int
	length           int
	delta            int
}

func main() {
	// part1()
	// no edits: 25s
	// toDestination param pointers: 29s
	part2()
	// fmt.Printf(math.MaxInt32) // 2147483647
}

func part1() {
	input := input1
	seeds, maps := parseInput(input)
	// fmt.Println("\nseeds", seeds)
	// fmt.Println("\nmaps[n]", maps[1])
	locs := make([]int, len(seeds))
	for i, seed := range seeds {
		// fmt.Println("seed: ", seed)
		source := seed
		for _, ranges := range maps {
			source = toDestination(&source, &ranges)
			// fmt.Println("new source: ", source)
		}
		locs[i] = source
	}
	fmt.Println("min loc", slices.Min(locs))
}

func part2() {
	input := input1
	seedPairs, maps := parseInput(input)
	seedStarts, seedLengths := parseSeedPairs(seedPairs)
	// 183212530, 346168560, 126471773, 4917124, 1095439147
	// fmt.Println(len(seedLengths))
	seedStarts = seedStarts[8:10]
	seedLengths = seedLengths[8:10]

	// generate seeds
	totalLength := 0
	for _, length := range seedLengths {
		totalLength += length
	}
	fmt.Println("total size in MB", totalLength*8/1000000)
	// return
	// allocate the entire size of the array to prevent many reallocations
	seeds := make([]int, totalLength)
	nextIndex := 0
	for i, start := range seedStarts {
		for seed := start; seed < start+seedLengths[i]; seed++ {
			seeds[nextIndex] = seed
			nextIndex++
		}
	}
	// fmt.Println("\nseeds", seeds)

	// find locs
	minLoc := math.MaxInt
	length := len(seeds)
	for i := 0; i < length; i++ {
		source := seeds[i]
		for j := 0; j < len(maps); j++ {
			source = toDestination(&source, &maps[j])
		}
		if source < minLoc {
			minLoc = source
		}
	}
	fmt.Println("minLoc: ", minLoc)
}

func parseInput(input string) ([]int, [][]Range) {
	groups := strings.Split(input, "\n\n")
	// fmt.Println("groups", groups[1])
	seedStrings := strings.Fields(strings.Split(groups[0], ": ")[1])
	seeds := make([]int, len(seedStrings))
	for i, seedStr := range seedStrings {
		num, _ := strconv.Atoi(seedStr)
		seeds[i] = num
	}
	mapGroups := groups[1:]
	ranges := make([][]Range, len(mapGroups))
	for i, group := range mapGroups {
		lines := strings.Split(group, "\n")[1:]
		for _, line := range lines {
			nums := stringsToInts(strings.Fields(line))
			var r Range
			r.destinationStart = nums[0]
			r.sourceStart = nums[1]
			r.length = nums[2]
			r.destinationEnd = r.destinationStart + r.length - 1
			r.sourceEnd = r.sourceStart + r.length - 1
			r.delta = r.destinationStart - r.sourceStart
			ranges[i] = append(ranges[i], r)
		}
	}
	return seeds, ranges
}

func parseSeedPairs(seedPairs []int) ([]int, []int) {
	seedStarts := make([]int, len(seedPairs)/2)
	seedLengths := make([]int, len(seedPairs)/2)
	for i := 0; i < len(seedPairs); i += 2 {
		seedStarts[i/2] = seedPairs[i]
		seedLengths[i/2] = seedPairs[i+1]
	}
	return seedStarts, seedLengths
}

func stringsToInts(strs []string) []int {
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

func toDestination(source *int, ranges *[]Range) int {
	for _, r := range *ranges {
		if *source >= r.sourceStart && *source <= r.sourceEnd {
			return *source + r.delta
		}
	}
	return *source
}
