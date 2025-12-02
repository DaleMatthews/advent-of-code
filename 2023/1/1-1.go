package main

import (
	"fmt"
	"strconv"
)

var input = []string{"fixme"}
var digitStrings = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	sum := 0
	for _, str := range input {
		first := 10
		last := 10
		for _, c := range str {
			num, err := strconv.Atoi(string(c))
			if err == nil {
				if first == 10 {
					first = num
				}
				last = num
			}
		}
		combined := first*10 + last
		sum += combined
		fmt.Println(combined)
	}
	fmt.Println(sum)
}
