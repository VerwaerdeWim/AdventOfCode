package main

import (
	"bytes"
	"strconv"
	"testing"
)

var testInput = []byte(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`)

func TestFn1(t *testing.T) {
	t.Run("Fn1", func(t *testing.T) {
		want := 2
		got, err := fn1(testInput)
		if err != nil {
			t.Fail()
		}
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	})
}

func TestIsReportSafe(t *testing.T) {
	testInputLines := bytes.Split(testInput, []byte("\n"))
	testLines := testInputLines[:len(testInputLines)-1]

	reports := make([][]int, 0, len(testInputLines))

	for _, testInputLine := range testLines {
		report := make([]int, 0, len(testInputLine))
		for _, level := range bytes.Split(testInputLine, []byte(" ")) {
			levelValue, err := strconv.Atoi(string(level))
			if err != nil {
				t.Fail()
			}
			report = append(report, levelValue)
		}
		reports = append(reports, report)
	}

	tests := []struct {
		input []int
		want  bool
	}{
		{
			input: reports[0],
			want:  true,
		},
		{
			input: reports[1],
			want:  false,
		},
		{
			input: reports[2],
			want:  false,
		},
		{
			input: reports[3],
			want:  false,
		},
		{
			input: reports[4],
			want:  false,
		},
		{
			input: reports[5],
			want:  true,
		},
		{
			input: []int{10, 6, 3, 2, 1},
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run("isReportSafe", func(t *testing.T) {
			got := isReportSafe(test.input)
			if got != test.want {
				t.Errorf("%d: Expected %t, got %t", test.input, test.want, got)
			}
		})
	}
}

func TestIsReportSafeWithDampener(t *testing.T) {
	testInputLines := bytes.Split(testInput, []byte("\n"))
	testLines := testInputLines[:len(testInputLines)-1]

	reports := make([][]int, 0, len(testInputLines))

	for _, testInputLine := range testLines {
		report := make([]int, 0, len(testInputLine))
		for _, level := range bytes.Split(testInputLine, []byte(" ")) {
			levelValue, err := strconv.Atoi(string(level))
			if err != nil {
				t.Fail()
			}
			report = append(report, levelValue)
		}
		reports = append(reports, report)
	}

	tests := []struct {
		input []int
		want  bool
	}{
		{
			input: reports[0],
			want:  true,
		},
		{
			input: reports[1],
			want:  false,
		},
		{
			input: reports[2],
			want:  false,
		},
		{
			input: reports[3],
			want:  true,
		},
		{
			input: reports[4],
			want:  true,
		},
		{
			input: reports[5],
			want:  true,
		},
		{
			input: []int{10, 6, 3, 2, 1},
			want:  true,
		},
		{
			input: []int{10, 10, 9, 8, 7, 5, 3},
			want:  true,
		},
		{
			input: []int{10, 6, 7, 5, 4, 2},
			want:  true,
		},
		{
			input: []int{7, 6, 2, 5},
			want:  true,
		},
		{
			input: []int{7, 6, 4, 0},
			want:  true,
		},
		{
			input: []int{10, 6, 2, 1},
			want:  false,
		},
		{
			input: []int{10, 6, 11, 12},
			want:  true,
		},
		{
			input: []int{10, 6, 7, 6},
			want:  true,
		},
		{
			input: []int{2, 5, 4, 3, 2},
			want:  true,
		},
	}
	for _, test := range tests {
		t.Run("isReportSafe", func(t *testing.T) {
			got := isReportSafeWithDampener(test.input)
			if got != test.want {
				t.Errorf("%d: Expected %t, got %t", test.input, test.want, got)
			}
		})
	}
}

func TestFn2(t *testing.T) {
	t.Run("Fn2", func(t *testing.T) {
		want := 4
		got, err := fn2(testInput)
		if err != nil {
			t.Fail()
		}
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	})
}
