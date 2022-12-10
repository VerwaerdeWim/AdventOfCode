package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	foodByElf := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	caloriesByElf := make([]int, len(foodByElf))

	calculateCaloriesByElf(foodByElf, caloriesByElf)

	sort.Sort(sort.Reverse(sort.IntSlice(caloriesByElf)))

	fmt.Println(caloriesByElf[0])
	fmt.Println(caloriesByElf[0] + caloriesByElf[1] + caloriesByElf[2])
}

func calculateCaloriesByElf(foodByElf []string, caloriesByElf []int) {
	for i, food := range foodByElf {
		for _, cals := range strings.Split(food, "\n") {
			calories, err := strconv.Atoi(cals)
			if err != nil {
				panic(err)
			}
			caloriesByElf[i] += calories
		}
	}
}
