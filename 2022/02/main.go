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
		other, me := normalizeInput(row)
		total += calcPoints(other, me)
	}

	fmt.Println(total)
}

// Rock  	-> 0
// Paper 	-> 1
// Scissors	-> 2
func normalizeInput(input string) (byte, byte) {
	other := input[0] - 65
	me := input[2] - 65 - 23
	return other, me
}

// Lose: 	(me-other+1+3)%3) = 0
// Draw: 	(me-other+1+3)%3) = 1
// Win: 	(me-other+1+3)%3) = 2
func calcPoints(other byte, me byte) int {
	return int(((me-other+1+3)%3)*3 + me + 1)
}

// X lose
// Y draw
// Z win
func part2(input []string) {
	var total int

	for _, row := range input {
		other, me := normalizeInput(row)
		updateMe(other, &me)
		total += calcPoints(other, me)
	}

	fmt.Println(total)
}

// Win:		me - 1 = 1
// Draw:	me - 1 = 0
// Lose:	me - 1 = -1
func updateMe(other byte, me *byte) {
	*me = (other + *me - 1 + 3) % 3
}
