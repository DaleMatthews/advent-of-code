package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var inputList1 = []string{"fixme"}

var inputList2 = []string{"fixme"}

func main() {
	inputList := inputList1
	sum := 0

	// create lists of indices of symbols for all lines
	symbolIndicies := getSymbolIndices(inputList)

	for i, input := range inputList {
		// grab relevant symbols
		startIndex := int(math.Max(float64(i)-1, 0))
		endIndex := int(math.Min(float64(i)+2, float64(len(inputList))))
		neighboringSymbols := symbolIndicies[startIndex:endIndex]
		sum += getLineSum(input, neighboringSymbols)
	}
	fmt.Println("Sum: ", sum)
}

func getSymbolIndices(inputList []string) [][]int {
	indicies := make([][]int, len(inputList))
	// have to escape the -
	reSymbol := regexp.MustCompile(`([!@#$%^&*()\-+=/])`)
	for i, input := range inputList {
		matches := reSymbol.FindAllStringIndex(input, -1)
		for _, match := range matches {
			indicies[i] = append(indicies[i], match[0])
		}
	}

	return indicies
}

var reNum = regexp.MustCompile(`([\d]+)`)

func getLineSum(input string, neighboringSymbols [][]int) int {
	// fmt.Println("input", input)
	// fmt.Println("neighboringSymbols", neighboringSymbols)
	lineSum := 0
	// capture number indices
	matches := reNum.FindAllStringIndex(input, -1)
	for _, match := range matches {
		numberStart := match[0]
		numberEnd := match[1] - 1
		isEnginePart := false
		// check if the start/end index neighbor any symbol
		for _, indices := range neighboringSymbols {
			for _, index := range indices {
				if numberStart < index && index < numberEnd {
					isEnginePart = true
				} else if math.Abs(float64(numberStart)-float64(index)) < 2 {
					isEnginePart = true
				} else if math.Abs(float64(numberEnd)-float64(index)) < 2 {
					isEnginePart = true
				}
			}
		}
		if isEnginePart {
			// parse number and add to sum
			numStr := input[numberStart : numberEnd+1]
			num, err := strconv.Atoi(numStr)
			if err == nil {
				lineSum += num
			}
		}
	}
	return lineSum
}
