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

func fn1(input []byte) (int, error) {
	var result int
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		levels := bytes.Split(line, []byte(" "))
		report := make([]int, 0, len(levels))
		for _, level := range levels {
			levelValue, err := strconv.Atoi(string(level))
			if err != nil {
				return result, err
			}
			report = append(report, levelValue)
		}
		if isReportSafe(report) {
			result++
		}
	}
	return result, nil
}

func isReportSafe(report []int) bool {
	desc := report[1] < report[0]

	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		if !checkDiff(diff, desc) {
			return false
		}
	}
	return true
}

func checkDiff(diff int, desc bool) bool {
	if diff == 0 {
		return false
	}
	if diff > 3 || diff < -3 {
		return false
	}
	if desc && diff < 0 || !desc && diff > 0 {
		return false
	}
	return true
}

func fn2(input []byte) (int, error) {
	var result int
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		levels := bytes.Split(line, []byte(" "))
		report := make([]int, 0, len(levels))
		for _, level := range levels {
			levelValue, err := strconv.Atoi(string(level))
			if err != nil {
				return result, err
			}
			report = append(report, levelValue)
		}
		if isReportSafeWithDampener(report) {
			result++
		}
	}
	return result, nil
}

// todo: make efficient
func isReportSafeWithDampener(report []int) bool {
	if !isReportSafe(report) {
		for i := range report {
			tempReport := make([]int, len(report))
			copy(tempReport, report)
			tempReport = slices.Delete(tempReport, i, i+1)
			if isReportSafe(tempReport) {
				return true
			}
		}
	} else {
		return true
	}
	return false
}
