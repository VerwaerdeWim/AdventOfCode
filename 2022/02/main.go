package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	part1(file)
	file.Seek(0, io.SeekStart)
	part2(file)
}

func part1(input io.Reader) {
	var total int

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		row := scanner.Text()
		other, me := normalizeInput(row)
		total += calcPoints(other, me)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
func part2(input io.Reader) {
	var total int

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		row := scanner.Text()
		other, me := normalizeInput(row)
		me = updateMe(other, me)
		total += calcPoints(other, me)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

// Win:		me - 1 = 1
// Draw:	me - 1 = 0
// Lose:	me - 1 = -1
func updateMe(other byte, me byte) byte {
	return (other + me - 1 + 3) % 3
}
