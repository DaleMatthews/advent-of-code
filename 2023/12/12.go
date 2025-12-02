package main

import (
	"1/util"
	"fmt"
	"strconv"
	"strings"
)

var input1 string = `fixme`

var input2 string = `fixme`

var cacheHit = 0
var cacheMiss = 0

func main() {
	input := input1
	data, metadata := parseInput(input)
	// util.PrintStringArray(data)
	// util.PrintBoard(metadata, " ")
	part1Sum := findPossibleMatches(data, metadata)
	for i := range data {
		data[i] = strings.Join([]string{data[i], data[i], data[i], data[i], data[i]}, "?")
		orig := metadata[i][:]
		metadata[i] = append(metadata[i], orig...)
		metadata[i] = append(metadata[i], orig...)
		metadata[i] = append(metadata[i], orig...)
		metadata[i] = append(metadata[i], orig...)
	}
	part2Sum := findPossibleMatches(data, metadata)
	// fmt.Println("Cache hits: ", cacheHit)
	// fmt.Println("Cache misses: ", cacheMiss)
	fmt.Println("part1Sum: ", part1Sum)
	fmt.Println("part2Sum: ", part2Sum)
}

func parseInput(input string) ([]string, [][]string) {
	lines := strings.Split(input, "\n")
	data := make([]string, len(lines))
	metadata := make([][]string, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		data[i] = parts[0]
		metadata[i] = strings.Split(parts[1], ",")
	}
	return data, metadata
}

func findPossibleMatches(data []string, metadata [][]string) int {
	// ???.### 1,1,3
	// .??..??...?##. 1,1,3
	// ?#?#?#?#?#?#?#? 1,3,1,6
	sum := 0
	for i, datum := range data {
		meta := metadata[i]
		sum += getPossibleCombinations(datum, meta)
		// fmt.Println("running total: ", sum)
	}
	return sum
}

var cache = make(map[string]int)

func getPossibleCombinations(datum string, meta []string) int {
	cacheKey := datum + strings.Join(meta, " ")
	value, hit := cache[cacheKey]
	if hit {
		cacheHit++
		return value
	}
	cacheMiss++

	if len(datum) == 0 {
		if len(meta) == 0 {
			return 1
		}
		return 0
	}
	if len(meta) == 0 {
		// no matchers left, make sure datum doesn't contain more # hits
		for _, c := range datum {
			if c == '#' {
				return 0
			}
		}
		return 1
	}

	if len(datum) < util.Sum(util.StringsToInts(meta))+len(meta)-1 {
		return 0 // not long enough
	}

	if datum[0] == '.' {
		value = getPossibleCombinations(datum[1:], meta)
		cache[cacheKey] = value
		return value
	}
	if datum[0] == '#' {
		count, _ := strconv.Atoi(meta[0])
		for i := 0; i < count; i++ {
			if i < len(datum) && datum[i] == '.' {
				value = 0 // not long enough
				cache[cacheKey] = value
				return value
			}
		}
		if count < len(datum) && datum[count] == '#' {
			value = 0 // too long
			cache[cacheKey] = value
			return value
		}

		nextDatum := ""
		if count+1 <= len(datum) {
			nextDatum = datum[count+1:]
		}
		nextMeta := make([]string, 0)
		if len(meta) > 0 {
			nextMeta = meta[1:]
		}
		value = getPossibleCombinations(nextDatum, nextMeta)
		cache[cacheKey] = value
		return value
	}
	variation1 := getPossibleCombinations("#"+datum[1:], meta)
	variation2 := getPossibleCombinations("."+datum[1:], meta)
	cache[cacheKey] = variation1 + variation2
	return cache[cacheKey]
}

// ! unused
// func buildRegexp(meta []string) *regexp.Regexp {
// 	str := `^\.*`
// 	for i, numStr := range meta {
// 		str += "#{" + numStr + "}"
// 		if i != len(meta)-1 {
// 			str += `\.+`
// 		}
// 	}
// 	str += `\.*$`
// 	// fmt.Println(str)
// 	return regexp.MustCompile(str)
// }

// ! couldn't get recursion to work with this code
// type getVariationsDef func(string, []string) int
// func memoize(f getVariationsDef) getVariationsDef {
// 	cache := make(map[string]int)
// 	return func(datum string, metadata []string) int {
// 		cacheKey := datum + strings.Join(metadata, " ")
// 		value, hit := cache[cacheKey]
// 		if hit {
// 			return value
// 		}
// 		value = f(datum, metadata)
// 		cache[cacheKey] = value
// 		return value
// 	}
// }
