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

type monkey struct {
	items     []int
	operation string
	divisible int
	test      map[bool]int
}

func main() {
	stringSlice := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	monkeys := make([]monkey, len(stringSlice))

	for i, m := range stringSlice {
		monkeys[i].test = make(map[bool]int, 2)

		for j, row := range strings.Split(strings.TrimRight(m, "\n"), "\n") {
			line := strings.TrimSpace(row)

			if j == 1 {
				startingItemsLine := strings.Split(line, ":")
				startingItems := strings.Split(startingItemsLine[1], ",")

				monkeys[i].items = make([]int, 0, len(startingItems))

				for _, startingItem := range startingItems {
					item, err := strconv.Atoi(strings.TrimSpace(startingItem))
					if err != nil {
						panic(err)
					}
					monkeys[i].items = append(monkeys[i].items, item)
				}
			}

			if j == 2 {
				operationLine := strings.Split(line, "= old")
				monkeys[i].operation = strings.TrimSpace(operationLine[1])
			}

			if j == 3 {
				testLine := strings.Split(line, "by")
				divisibleBy, err := strconv.Atoi(strings.TrimSpace(testLine[1]))
				if err != nil {
					panic(err)
				}
				monkeys[i].divisible = divisibleBy
			}
			if j == 4 {
				trueLine := strings.Split(line, "monkey")
				ifTrue, err := strconv.Atoi(strings.TrimSpace(trueLine[1]))
				if err != nil {
					panic(err)
				}
				monkeys[i].test[true] = ifTrue
			}
			if j == 5 {
				falseLine := strings.Split(line, "monkey")
				ifFalse, err := strconv.Atoi(strings.TrimSpace(falseLine[1]))
				if err != nil {
					panic(err)
				}
				monkeys[i].test[false] = ifFalse
			}
		}
	}

	monkeysCopy := make([]monkey, len(stringSlice))
	copy(monkeysCopy, monkeys)
	for i, monkey := range monkeys {
		monkeysCopy[i].items = make([]int, len(monkey.items))
		copy(monkeysCopy[i].items, monkey.items)

	}
	part1(monkeys)
	part2(monkeysCopy)
}

func part1(monkeys []monkey) {
	inspections := make([]int, len(monkeys))
	for round := 0; round < 20; round++ {
		for i := 0; i < len(monkeys); i++ {
			for j, item := range monkeys[i].items {
				operation := strings.Split(monkeys[i].operation, " ")
				operationValue := item
				if operation[1] != "old" {
					value, err := strconv.Atoi(operation[1])
					if err != nil {
						panic(err)
					}
					operationValue = value
				}

				switch operation[0] {
				case "*":
					monkeys[i].items[j] *= operationValue
				case "+":
					monkeys[i].items[j] += operationValue
				}
				monkeys[i].items[j] /= 3
				throwTo := monkeys[i].test[monkeys[i].items[j]%monkeys[i].divisible == 0]
				inspections[i]++
				monkeys[throwTo].items = append(monkeys[throwTo].items, monkeys[i].items[j])

			}
			monkeys[i].items = nil
		}

	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	fmt.Println(inspections[0] * inspections[1])
}

func part2(monkeys []monkey) {
	// take the modulo of the product of all divisors,
	// this will make sure that the result of the divisor check won't change

	// least common multiple is the product of all divisors as all divisors are prime numbers
	lcm := 1
	for _, monkey := range monkeys {
		lcm *= monkey.divisible
	}

	inspections := make([]int, len(monkeys))
	for round := 0; round < 10000; round++ {
		for i := 0; i < len(monkeys); i++ {
			for j, item := range monkeys[i].items {
				operation := strings.Split(monkeys[i].operation, " ")
				operationValue := item
				if operation[1] != "old" {
					value, err := strconv.Atoi(operation[1])
					if err != nil {
						panic(err)
					}
					operationValue = value
				}

				switch operation[0] {
				case "*":
					monkeys[i].items[j] *= operationValue
				case "+":
					monkeys[i].items[j] += operationValue
				}
				monkeys[i].items[j] %= lcm
				// monkeys[i].items[j] /= 3
				throwTo := monkeys[i].test[monkeys[i].items[j]%monkeys[i].divisible == 0]
				inspections[i]++
				monkeys[throwTo].items = append(monkeys[throwTo].items, monkeys[i].items[j])

			}
			monkeys[i].items = nil
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	fmt.Println(inspections[0] * inspections[1])
}
