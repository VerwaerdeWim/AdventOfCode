package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"slices"
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

func prepSection1(section1 [][]byte) (map[int][]int, error) {
	ruleMap := make(map[int][]int)
	for _, line := range section1 {
		digits := bytes.Split(line, []byte("|"))
		page1, err := strconv.Atoi(string(digits[0]))
		if err != nil {
			return nil, err
		}

		page2, err := strconv.Atoi(string(digits[1]))
		if err != nil {
			return nil, err
		}
		ruleMap[page1] = append(ruleMap[page1], page2)
	}
	return ruleMap, nil
}

func prepSection2(section2 [][]byte) ([][]int, error) {
	updates := make([][]int, 0, len(section2))
	for _, line := range section2 {
		digits := bytes.Split(line, []byte(","))
		update := make([]int, 0, len(digits))
		for _, digit := range digits {
			page, err := strconv.Atoi(string(digit))
			if err != nil {
				return nil, err
			}
			update = append(update, page)
		}
		updates = append(updates, update)
	}
	return updates, nil
}

func fn1(input []byte) (int, error) {
	var result int
	section1AndSection2 := bytes.Split(input, []byte("\n\n"))
	section1 := bytes.Split(section1AndSection2[0], []byte("\n"))
	section2 := bytes.Split(section1AndSection2[1], []byte("\n"))
	section2 = section2[:len(section2)-1]
	ruleMap, err := prepSection1(section1)
	if err != nil {
		return result, err
	}

	updates, err := prepSection2(section2)
	if err != nil {
		return result, err
	}

	for _, update := range updates {
		if checkUpdate(update, ruleMap) {
			result += update[(len(update)-1)/2]
		}
	}

	return result, nil
}

func checkUpdate(update []int, ruleMap map[int][]int) bool {
	for i := 0; i < len(update)-1; i++ {
		if _, exists := ruleMap[update[i]]; !exists {
			if slices.Contains(ruleMap[update[i+1]], update[i]) {
				return false
			}
		} else {
			if !slices.Contains(ruleMap[update[i]], update[i+1]) {
				return false
			}
		}

	}
	return true
}

func fn2(input []byte) (int, error) {
	var result int
	section1AndSection2 := bytes.Split(input, []byte("\n\n"))
	section1 := bytes.Split(section1AndSection2[0], []byte("\n"))
	section2 := bytes.Split(section1AndSection2[1], []byte("\n"))
	section2 = section2[:len(section2)-1]
	ruleMap, err := prepSection1(section1)
	if err != nil {
		return result, err
	}

	updates, err := prepSection2(section2)
	if err != nil {
		return result, err
	}

	for _, update := range updates {
		if !checkUpdate(update, ruleMap) {
			correctedUpdate := correctUpdate(update, ruleMap)
			result += correctedUpdate[(len(correctedUpdate)-1)/2]
		}
	}

	return result, nil
}

// todo: make efficient
func correctUpdate(update []int, ruleMap map[int][]int) []int {
	i := 0
	for i < len(update) {
		page := update[i]
		for _, rule := range ruleMap[page] {
			previousPages := update[:i]
			index := slices.Index(previousPages, rule)
			if index != -1 {
				update = slices.Delete(update, index, index+1)
				update = slices.Insert(update, i, rule)
				i = 0
				break
			}
		}
		i++
	}
	return update
}
