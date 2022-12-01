package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var list []int
	var tempNumber int
	var max int

	for scanner.Scan() {
		if scanner.Text() != "" {
			number, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			tempNumber += number
		} else {
			if tempNumber > max {
				max = tempNumber
			}
			list = append(list, tempNumber)
			tempNumber = 0
		}
	}

	if tempNumber > max {
		max = tempNumber
	}

	list = append(list, tempNumber)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(max)

	sort.Slice(list, func(i, j int) bool {
		return list[j] < list[i]
	})

	fmt.Println(list[0] + list[1] + list[2])
}
