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
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)

	var counter int
	var index int
	set := make(map[string]struct{}, 4)
	marker := make([]string, 0, 4)
	for scanner.Scan() {

		index++
		letter := scanner.Text()
		fmt.Print(letter)
		fmt.Println(counter)
		if _, present := set[letter]; present {
			for i, val := range marker {
				delete(set, val)
				counter--
				if val == letter {
					marker = marker[i+1:]
					break
				}
			}
			fmt.Println()
		}
		set[letter] = struct{}{}
		marker = append(marker, letter)

		counter++
		if counter == 4 {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Print(index)
}

func part2(input io.Reader) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)

	var counter int
	var index int
	set := make(map[string]struct{}, 14)
	marker := make([]string, 0, 14)
	for scanner.Scan() {

		index++
		letter := scanner.Text()
		fmt.Print(letter)
		fmt.Println(counter)
		if _, present := set[letter]; present {
			for i, val := range marker {
				delete(set, val)
				counter--
				if val == letter {
					marker = marker[i+1:]
					break
				}
			}
			fmt.Println()
		}
		set[letter] = struct{}{}
		marker = append(marker, letter)

		counter++
		if counter == 14 {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Print(index)
}
