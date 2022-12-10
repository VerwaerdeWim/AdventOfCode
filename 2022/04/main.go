package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	stringSlice := strings.Split(strings.TrimRight(input, "\n"), "\n")
	part1(stringSlice)
	part2(stringSlice)
}

func part1(input []string) {
	var total int
	for _, row := range input {
		var firstStart, firstEnd, secondStart, secondEnd int
		fmt.Sscanf(row, "%d-%d,%d-%d", &firstStart, &firstEnd, &secondStart, &secondEnd)
		if firstStart <= secondStart && firstEnd >= secondEnd {
			total++
		} else if secondStart <= firstStart && secondEnd >= firstEnd {
			total++
		}
	}
	fmt.Println(total)
}

func part2(input []string) {
	var total int
	for _, row := range input {
		var firstStart, firstEnd, secondStart, secondEnd int
		fmt.Sscanf(row, "%d-%d,%d-%d", &firstStart, &firstEnd, &secondStart, &secondEnd)
		if firstStart <= secondEnd && firstStart >= secondStart {
			total++
		} else if secondStart <= firstEnd && secondStart >= firstStart {
			total++
		}
	}
	fmt.Println(total)
}
