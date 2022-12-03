package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
		size := len(row) / 2
		comp1 := row[:size]
		comp2 := row[size:]
		var item rune
		for _, j := range comp1 {
			if strings.Contains(comp2, string(j)) {
				item = j
				break
			}
		}
		// A = 65 + 32 = 97 - 64 = 33 - 6 = 27
		// a = 97 - 32 = 65 - 64 = 1
		if item < 97 {
			item -= 38
		} else {
			item -= 96
		}
		total += int(item)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

func part2(input io.Reader) {
	scanner := bufio.NewScanner(input)
	var i int
	var sacks [3]string
	var total int
	for scanner.Scan() {
		row := scanner.Text()
		sacks[i] = row
		i = (i + 1) % 3
		if i == 0 {
			sack1 := sacks[0]
			sack2 := sacks[1]
			sack3 := sacks[2]

			intersection := map[rune]struct{}{}
			var badge rune
			for _, j := range sack1 {
				if strings.Contains(sack2, string(j)) {
					intersection[j] = struct{}{}
				}
			}
			for k := range intersection {
				if strings.Contains(sack3, string(k)) {
					badge = k
					break
				}
			}
			if badge < 97 {
				badge -= 38
			} else {
				badge -= 96
			}
			total += int(badge)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
