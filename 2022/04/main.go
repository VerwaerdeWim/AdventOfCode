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
		var firstStart, firstEnd, secondStart, secondEnd int
		fmt.Sscanf(row, "%d-%d,%d-%d", &firstStart, &firstEnd, &secondStart, &secondEnd)
		if firstStart <= secondStart && firstEnd >= secondEnd {
			total++
		} else if secondStart <= firstStart && secondEnd >= firstEnd {
			total++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}

func part2(input io.Reader) {
	var total int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := scanner.Text()
		var firstStart, firstEnd, secondStart, secondEnd int
		fmt.Sscanf(row, "%d-%d,%d-%d", &firstStart, &firstEnd, &secondStart, &secondEnd)
		if firstStart <= secondEnd && firstStart >= secondStart {
			total++
		} else if secondStart <= firstEnd && secondStart >= firstStart {
			total++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
