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

// todo: only works if the stacks amount are single digits
func part1(input io.Reader) {
	scanner := bufio.NewScanner(input)
	unorderedStacks := make(map[int][]rune)

	for scanner.Scan() {
		row := scanner.Text()
		for i, v := range row {
			if i%2 != 0 && v != ' ' {
				unorderedStacks[i] = append(unorderedStacks[i], v)
			}
		}
		if len(row) == 0 {
			break
		}

	}

	stacks := make(map[rune][]rune)
	for _, stack := range unorderedStacks {
		index := stack[len(stack)-1]
		for i := len(stack) - 2; i >= 0; i-- {
			stacks[index] = append(stacks[index], stack[i])
		}
	}

	for scanner.Scan() {

		row := scanner.Text()
		var amount int
		var from, to rune
		fmt.Sscanf(row, "move %d from %c to %c", &amount, &from, &to)

		toMove := stacks[from][len(stacks[from])-amount : len(stacks[from])]
		stacks[from] = stacks[from][:len(stacks[from])-amount]
		for i := len(toMove) - 1; i >= 0; i-- {
			stacks[to] = append(stacks[to], toMove[i])
		}
	}

	result := make([]rune, 9)
	for k, stack := range stacks {
		result[k-49] = stack[len(stack)-1]
	}
	fmt.Println(string(result))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part2(input io.Reader) {
	scanner := bufio.NewScanner(input)
	unorderedStacks := make(map[int][]rune)

	for scanner.Scan() {
		row := scanner.Text()
		for i, v := range row {
			if i%2 != 0 && v != ' ' {
				unorderedStacks[i] = append(unorderedStacks[i], v)
			}
		}
		if len(row) == 0 {
			break
		}

	}

	stacks := make(map[rune][]rune)
	for _, stack := range unorderedStacks {
		index := stack[len(stack)-1]
		for i := len(stack) - 2; i >= 0; i-- {
			stacks[index] = append(stacks[index], stack[i])
		}
	}

	for scanner.Scan() {

		row := scanner.Text()
		var amount int
		var from, to rune
		fmt.Sscanf(row, "move %d from %c to %c", &amount, &from, &to)

		toMove := stacks[from][len(stacks[from])-amount : len(stacks[from])]
		stacks[from] = stacks[from][:len(stacks[from])-amount]
		for i := 0; i < len(toMove); i++ {
			stacks[to] = append(stacks[to], toMove[i])
		}
	}

	result := make([]rune, 9)
	for k, stack := range stacks {
		result[k-49] = stack[len(stack)-1]
	}
	fmt.Println(string(result))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
