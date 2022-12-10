package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	stringSlice := strings.Split(strings.TrimRight(input, "\n"), "\n")
	part1(stringSlice)
	part2(stringSlice)
}

func part1(stringSlice []string) {
	cycle := 0
	x := 1
	result := 0
	for _, row := range stringSlice {
		cycle++
		result += checkSignalStrength(cycle, x)
		instruction := strings.Split(row, " ")
		if len(instruction) == 1 {
			continue
		}
		cycle++
		result += checkSignalStrength(cycle, x)
		value, err := strconv.Atoi(instruction[1])
		if err != nil {
			panic(err)
		}
		x += value
	}
	fmt.Println(result)

}

func part2(stringSlice []string) {
	cycle := 0
	x := 1
	for _, row := range stringSlice {
		cycle++
		drawCRT(&cycle, x)
		instruction := strings.Split(row, " ")
		if len(instruction) == 1 {
			continue
		}
		cycle++
		drawCRT(&cycle, x)
		value, err := strconv.Atoi(instruction[1])
		if err != nil {
			panic(err)
		}
		x += value
	}
}

func checkSignalStrength(cycle int, x int) int {

	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		return cycle * x
	}

	return 0
}

func drawCRT(cycle *int, x int) {
	if *cycle <= x+1+1 && *cycle >= x+1-1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if *cycle == 40 || *cycle == 80 || *cycle == 120 || *cycle == 160 || *cycle == 200 || *cycle == 240 {
		*cycle -= 40
		fmt.Println()
	}
}
