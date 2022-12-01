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

	for scanner.Scan() {

		if scanner.Text() != "" {
			number, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			tempNumber += number
		} else {
			list = append(list, tempNumber)
			tempNumber = 0
		}
	}
	list = append(list, tempNumber)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var max int
	for _, value := range list {
		if value > max {
			max = value
		}
	}

	fmt.Println(max)

	sort.Slice(list, func(i, j int) bool {
		return list[j] < list[i]
	})

	fmt.Println(list[0] + list[1] + list[2])

}
