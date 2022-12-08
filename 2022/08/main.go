package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	stringSlice := strings.Split(strings.TrimRight(input, "\n"), "\n")
	part1(stringSlice)
	part2(stringSlice)

}

func part1(input []string) {
	heightMap := make([][]int, len(input))
	visibleMap := make([][]bool, len(input))
	for i, row := range input {
		heightMap[i] = make([]int, len(row))
		visibleMap[i] = make([]bool, len(row))
		for j, char := range row {
			heightMap[i][j] = int(char - '0')
		}
	}

	var result int

	length := len(heightMap)
	// check from north to south and from west to east
	for i := 0; i < length; i++ {
		k := length - 1 - i // k decrements from colLength-1 to 0
		rowMaxLeft := heightMap[i][0]
		colMaxTop := heightMap[0][i]

		rowMaxRight := heightMap[length-1-i][length-1]
		colMaxBottom := heightMap[length-1][length-1-i]
		for j := 0; j < length; j++ {
			l := length - 1 - j //l decrements from len(rowLength-1 to 0)
			// north side and south side
			if i == 0 || i == length-1 {
				result += setVisible(&visibleMap[i][j])
				continue
			}
			// west side and east side
			if j == 0 || j == length-1 {
				result += setVisible(&visibleMap[i][j])
				continue
			}

			// from west to east
			if heightMap[i][j] > rowMaxLeft {
				rowMaxLeft = heightMap[i][j]
				result += setVisible(&visibleMap[i][j])
			}

			// from north to south
			if heightMap[j][i] > colMaxTop {
				colMaxTop = heightMap[j][i]
				result += setVisible(&visibleMap[j][i])
			}

			// from east to west
			if heightMap[k][l] > rowMaxRight {
				rowMaxRight = heightMap[k][l]
				result += setVisible(&visibleMap[k][l])
			}

			// from south to north
			if heightMap[l][k] > colMaxBottom {
				colMaxBottom = heightMap[l][k]
				result += setVisible(&visibleMap[l][k])
			}
		}
	}
	// printMap(visibleMap, heightMap)
	fmt.Println(result)
}

func setVisible(visible *bool) int {
	if !*visible {
		*visible = true
		return 1
	}
	return 0
}

func printMap(visibleMap [][]bool, heightMap [][]int) {
	for i, row := range visibleMap {
		for j, visible := range row {
			if visible {
				fmt.Print(heightMap[i][j])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func part2(input []string) {
	heightMap := make([][]int, len(input))
	scoreMap := make([][]int, len(input))

	for i, row := range input {
		heightMap[i] = make([]int, len(row))
		scoreMap[i] = make([]int, len(row))
		for j, char := range row {
			heightMap[i][j] = int(char - '0')
			scoreMap[i][j] = 1
		}
	}

	length := len(heightMap)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			// north side and south side
			if i == 0 || i == length-1 {
				scoreMap[i][j] = 0
				continue
			}
			// west side and east side
			if j == 0 || j == length-1 {
				scoreMap[i][j] = 0
				continue
			}
			// north
			distance := 0
			for k := i - 1; k >= 0; k-- {
				if heightMap[k][j] >= heightMap[i][j] {
					distance++
					break
				}
				distance++
			}
			scoreMap[i][j] *= distance

			// south
			distance = 0
			for k := i + 1; k < length; k++ {
				if heightMap[k][j] >= heightMap[i][j] {
					distance++
					break
				}
				distance++
			}
			scoreMap[i][j] *= distance

			// east
			distance = 0
			for k := j + 1; k < length; k++ {
				if heightMap[i][k] >= heightMap[i][j] {
					distance++
					break
				}
				distance++
			}
			scoreMap[i][j] *= distance

			// west
			distance = 0
			for k := j - 1; k >= 0; k-- {
				if heightMap[i][k] >= heightMap[i][j] {
					distance++
					break
				}
				distance++
			}
			scoreMap[i][j] *= distance
		}
	}

	result := 0
	for _, row := range scoreMap {
		for _, score := range row {
			if score > result {
				result = score
			}
		}
	}
	fmt.Println(result)
}
