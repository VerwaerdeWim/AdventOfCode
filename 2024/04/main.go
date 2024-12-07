package main

import (
	"bytes"
	_ "embed"
	"fmt"
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

type Direction struct {
	dx int
	dy int
}

func fn1(input []byte) (int, error) {
	var result int
	word := []byte("XMAS")
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]

	directions := []Direction{
		{-1, 0}, {1, 0},
		{0, -1}, {0, 1},
		{-1, -1}, {1, -1},
		{-1, 1}, {1, 1},
	}

	for y, line := range lines {
		for x := range line {
			for _, dir := range directions {
				if checkDirection(lines, x, y, dir.dx, dir.dy, word) {
					result++
				}
			}
		}
	}
	return result, nil
}

func checkDirection(lines [][]byte, x, y, dx, dy int, word []byte) bool {
	for i, char := range word {
		nx, ny := x+i*dx, y+i*dy

		if ny < 0 || ny >= len(lines) || nx < 0 || nx >= len(lines[ny]) || lines[ny][nx] != char {
			return false
		}
	}

	return true
}

func fn2(input []byte) (int, error) {
	var result int
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]
	for y, line := range lines {
		for x, char := range line {
			if y >= 1 && x >= 1 && y < len(lines)-1 && x < len(line)-1 {
				if char == 'A' {
					if ((lines[y+1][x-1] == 'M' && lines[y-1][x+1] == 'S') || (lines[y+1][x-1] == 'S' && lines[y-1][x+1] == 'M')) &&
						((lines[y+1][x+1] == 'M' && lines[y-1][x-1] == 'S') || (lines[y+1][x+1] == 'S' && lines[y-1][x-1] == 'M')) {
						result++
					}
				}
			}
		}
	}
	return result, nil
}
