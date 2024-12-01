package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

//go:embed input.txt
var input []byte

func main() {
	part1, err := fn1(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result part 1:%d\n", part1)

	part2, err := fn2(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result part 2:%d\n", part2)
}

func prepData(input []byte) ([]int, []int, error) {
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]

	list1 := make([]int, 0, len(lines))
	list2 := make([]int, 0, len(lines))
	for _, line := range lines {
		re := regexp.MustCompile(`(\d+)\s{3}(\d+)`)
		matches := re.FindSubmatch(line)

		id1, err := strconv.Atoi(string(matches[1]))
		if err != nil {
			return list1, list2, err
		}
		list1 = append(list1, id1)

		id2, err := strconv.Atoi(string(matches[2]))
		if err != nil {
			return list1, list2, err
		}
		list2 = append(list2, id2)
	}
	return list1, list2, nil
}

func fn1(input []byte) (int, error) {
	var result int
	list1, list2, err := prepData(input)
	if err != nil {
		return result, err
	}
	result = calculateDistance(list1, list2)
	return result, nil
}

func calculateDistance(col1 []int, col2 []int) int {
	var totalDistance int
	sort.Ints(col1)
	sort.Ints(col2)

	for i, c1 := range col1 {
		distance := c1 - col2[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}
	return totalDistance
}

func fn2(input []byte) (int, error) {
	var result int

	list1, list2, err := prepData(input)
	if err != nil {
		return result, err
	}

	result = calculateSimilarity(list1, list2)
	return result, nil
}

func calculateSimilarity(list1 []int, list2 []int) int {
	var similarity int

	freqMap := make(map[int]int)
	for _, c2 := range list2 {
		freqMap[c2]++
	}

	for _, c1 := range list1 {
		similarity += c1 * freqMap[c1]
	}
	return similarity
}
