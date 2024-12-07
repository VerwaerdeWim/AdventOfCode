package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"regexp"
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

func fn1(input []byte) (int, error) {
	var result int
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]
	for _, line := range lines {
		res, err := scanLine1(line)
		if err != nil {
			return result, err
		}
		result += res
	}
	return result, nil
}

func scanLine1(line []byte) (int, error) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllSubmatch(line, -1)
	if matches == nil {
		return 0, nil
	}

	var result int
	for _, match := range matches {
		res, err := multiply(match[1], match[2])
		if err != nil {
			return result, err
		}
		result += res
	}

	return result, nil
}

func multiply(digit1 []byte, digit2 []byte) (int, error) {
	x, err := strconv.Atoi(string(digit1))
	if err != nil {
		return 0, err
	}

	y, err := strconv.Atoi(string(digit2))
	if err != nil {
		return 0, err
	}
	return x * y, nil
}

func fn2(input []byte) (int, error) {
	var result int
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]
	enabled := true
	for _, line := range lines {
		var res int
		var err error
		res, enabled, err = scanLine2(line, enabled)
		if err != nil {
			return result, err
		}
		result += res
	}
	return result, nil
}

type Status int

const (
	START Status = iota + 1
	DIGIT1
	DIGIT2
)

func scanLine2(line []byte, enabled bool) (int, bool, error) {
	var result int
	status := START
	mul := []byte("mul(")
	dont := []byte("don't()")
	do := []byte("do()")
	digit1 := make([]byte, 0, 3)
	digit2 := make([]byte, 0, 3)
	i := 0
	for i < len(line) {
		switch status {
		case START:
			if enabled {
				if jump := checkWindow(i, line, mul); jump > 0 {
					status = DIGIT1
					i += jump
					continue
				}
				if jump := checkWindow(i, line, dont); jump > 0 {
					enabled = false
					i += jump
					continue
				}
			} else if jump := checkWindow(i, line, do); jump > 0 {
				enabled = true
				i += jump
				continue
			}
			i++
		case DIGIT1:
			if validDigit(line[i], len(digit1)) {
				digit1 = append(digit1, line[i])
			} else if line[i] == ',' {
				status = DIGIT2
			} else {
				resetStatus(&status, &digit1, &digit2)
			}
			i++
		case DIGIT2:
			if validDigit(line[i], len(digit2)) {
				digit2 = append(digit2, line[i])
			} else if line[i] == ')' {
				res, err := multiply(digit1, digit2)
				if err != nil {
					return result, enabled, err
				}
				result += res
				resetStatus(&status, &digit1, &digit2)
			} else {
				resetStatus(&status, &digit1, &digit2)
			}
			i++
		default:
			panic("unhandled default case")
		}
	}
	return result, enabled, nil
}

func checkWindow(i int, line []byte, pattern []byte) int {
	for j := 0; j < len(pattern) && i+j < len(line); j++ {
		if i+j >= len(line) || line[i+j] != pattern[j] {
			return 0
		}
	}
	return len(pattern)
}

func validDigit(char byte, digitLen int) bool {
	if (digitLen == 0 && char >= '1' && char <= '9') ||
		(digitLen < 3 && char >= '0' && char <= '9') {
		return true
	}
	return false
}

func resetStatus(status *Status, digit1 *[]byte, digit2 *[]byte) {
	*status = START
	*digit1 = (*digit1)[:0]
	*digit2 = (*digit2)[:0]
}
