package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var input = []string{"fixme"}
var digitStrings = []*regexp.Regexp{
	regexp.MustCompile("one"),
	regexp.MustCompile("two"),
	regexp.MustCompile("three"),
	regexp.MustCompile("four"),
	regexp.MustCompile("five"),
	regexp.MustCompile("six"),
	regexp.MustCompile("seven"),
	regexp.MustCompile("eight"),
	regexp.MustCompile("nine"),
}

func main() {
	sum := 0
	for _, str := range input {
		first := 10
		last := 10
		for i, c := range str {
			num, err := strconv.Atoi(string(c))
			// if the numerical digit exists, use that
			if err == nil {
				if first == 10 {
					first = num
				}
				last = num
			} else {
				// otherwise, check for a valid digit word at this index
				partial := str[i:]
				for j, re := range digitStrings {
					loc := re.FindStringIndex(partial)
					if len(loc) > 0 && loc[0] == 0 {
						if first == 10 {
							first = j + 1
						}
						last = j + 1
					}
				}
			}
		}
		combined := first*10 + last
		sum += combined
		fmt.Println(combined)
	}
	fmt.Println(sum)
}
