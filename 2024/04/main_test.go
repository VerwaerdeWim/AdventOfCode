package main

import "testing"

var testInput1 = []byte(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`)

var testInput2 = []byte(`
`)

func TestFn1(t *testing.T) {
	t.Run("Fn1", func(t *testing.T) {
		want := 18
		got, err := fn1(testInput1)
		if err != nil {
			t.Fail()
		}
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	})
}

func TestFn2(t *testing.T) {
	t.Run("Fn2", func(t *testing.T) {
		want := 9
		got, err := fn2(testInput1)
		if err != nil {
			t.Fail()
		}
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	})
}
