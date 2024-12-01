package main

import "testing"

var testInput1 = []byte(`3   4
4   3
2   5
1   3
3   9
3   3
`)

func TestFn1(t *testing.T) {
	t.Run("Fn1", func(t *testing.T) {
		want := 11
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
		want := 31
		got, err := fn2(testInput1)
		if err != nil {
			t.Fail()
		}
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	})
}
