package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var inputList1 = []string{"fixme"}

var inputList2 = []string{"fixme"}

type Part struct {
	startIndex int
	endIndex   int
	value      int
}

func main() {
	inputList := inputList1
	gearRatioSum := 0

	// create lists of indices of symbols for all lines
	symbolIndicies := getSymbolIndices(inputList)
	parts := getParts(inputList)
	// fmt.Println("parts", parts)

	for i, symbols := range symbolIndicies {
		// grab relevant parts
		startIndex := int(math.Max(float64(i)-1, 0))
		endIndex := int(math.Min(float64(i)+2, float64(len(inputList))))
		neighboringParts := parts[startIndex:endIndex]
		gearRatioSum += getGearRatioLineSum(symbols, neighboringParts)
	}
	fmt.Println("gearRatioSum: ", gearRatioSum)
}

func getSymbolIndices(inputList []string) [][]int {
	indicies := make([][]int, len(inputList))
	reSymbol := regexp.MustCompile(`([*])`)
	for i, input := range inputList {
		matches := reSymbol.FindAllStringIndex(input, -1)
		for _, match := range matches {
			indicies[i] = append(indicies[i], match[0])
		}
	}

	return indicies
}

func getParts(inputList []string) [][]Part {
	parts := make([][]Part, len(inputList))
	var reNum = regexp.MustCompile(`([\d]+)`)
	for i, input := range inputList {
		matches := reNum.FindAllStringIndex(input, -1)
		for _, match := range matches {
			startIndex := match[0]
			endIndex := match[1] - 1
			numStr := input[startIndex : endIndex+1]
			value, _ := strconv.Atoi(numStr)
			parts[i] = append(parts[i], Part{startIndex, endIndex, value})
		}
	}

	return parts
}

func getGearRatioLineSum(symbols []int, neighboringParts [][]Part) int {
	// fmt.Println("symbols", symbols)
	// fmt.Println("neighboringParts", neighboringParts)
	lineSum := 0
	for _, symbolIndex := range symbols {
		neighbors := make([]int, 0, 2)
		// check if the start/end index neighbor any symbol
		for _, parts := range neighboringParts {
			for _, part := range parts {
				if part.startIndex < symbolIndex && symbolIndex < part.endIndex {
					neighbors = append(neighbors, part.value)
				} else if math.Abs(float64(part.startIndex)-float64(symbolIndex)) < 2 {
					neighbors = append(neighbors, part.value)
				} else if math.Abs(float64(part.endIndex)-float64(symbolIndex)) < 2 {
					neighbors = append(neighbors, part.value)
				}
			}
		}
		// fmt.Println("neighbors", neighbors)
		if len(neighbors) == 2 {
			lineSum += neighbors[0] * neighbors[1]
		}
	}
	return lineSum
}
