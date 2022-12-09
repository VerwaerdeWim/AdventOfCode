package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type knot struct {
	x int
	y int
}

func main() {
	stringSlice := strings.Split(strings.TrimRight(input, "\n"), "\n")
	fmt.Println(solve(2, stringSlice))
	fmt.Println(solve(10, stringSlice))
}

func solve(size int, input []string) int {
	knots := make([]knot, size)

	coordSet := make(map[int]map[int]struct{})
	for _, row := range input {
		instruction := strings.Split(row, " ")
		direction := instruction[0]
		steps, err := strconv.Atoi(instruction[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "U":
			for i := 0; i < steps; i++ {
				knots[0].y++
				moveRope(knots, coordSet)
			}
		case "R":
			for i := 0; i < steps; i++ {
				knots[0].x++
				moveRope(knots, coordSet)
			}
		case "D":
			for i := 0; i < steps; i++ {
				knots[0].y--
				moveRope(knots, coordSet)
			}
		case "L":
			for i := 0; i < steps; i++ {
				knots[0].x--
				moveRope(knots, coordSet)
			}
		}

	}
	var result int
	for _, row := range coordSet {
		result += len(row)
	}
	return result
}

func moveRope(knots []knot, locations map[int]map[int]struct{}) {
	for i := 1; i < len(knots); i++ {
		head := &knots[i-1]
		tail := &knots[i]
		move(head, tail, locations)
	}

	lastKnot := knots[len(knots)-1]
	if _, exist := locations[lastKnot.x]; !exist {
		locations[lastKnot.x] = make(map[int]struct{})
	}

	locations[lastKnot.x][lastKnot.y] = struct{}{}
}

func move(head *knot, tail *knot, locations map[int]map[int]struct{}) {
	deltaX := math.Abs(float64(head.x - tail.x))
	deltaY := math.Abs(float64(head.y - tail.y))
	if deltaX <= 1 && deltaY <= 1 {
		return
	}

	xDir := 1
	if head.x < tail.x {
		xDir = -1
	}
	yDir := 1
	if head.y < tail.y {
		yDir = -1
	}

	if head.x == tail.x {
		tail.y += yDir
		return
	}

	if head.y == tail.y {
		tail.x += xDir
		return
	}

	tail.x += xDir
	tail.y += yDir
}
