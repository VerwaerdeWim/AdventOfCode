package main

import (
	"bytes"
	"testing"
)

var testInput = []byte(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`)

func TestFn1(t *testing.T) {
	t.Run("Fn1", func(t *testing.T) {
		want := 143
		got, err := fn1(testInput)
		if err != nil {
			t.Fail()
		}
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	})
}

func TestCheckUpdate(t *testing.T) {
	section1AndSection2 := bytes.Split(testInput, []byte("\n\n"))
	section1 := bytes.Split(section1AndSection2[0], []byte("\n"))
	section2 := bytes.Split(section1AndSection2[1], []byte("\n"))
	section2 = section2[:len(section2)-1]

	ruleMap, err := prepSection1(section1)
	if err != nil {
		t.Fail()
	}

	updates, err := prepSection2(section2)
	if err != nil {
		t.Fail()
	}

	tests := []struct {
		ruleMap map[int][]int
		update  []int
		want    bool
	}{
		{
			ruleMap: ruleMap,
			update:  updates[0],
			want:    true,
		},
		{
			ruleMap: ruleMap,
			update:  updates[1],
			want:    true,
		},
		{
			ruleMap: ruleMap,
			update:  updates[2],
			want:    true,
		},
		{
			ruleMap: ruleMap,
			update:  updates[3],
			want:    false,
		},
		{
			ruleMap: ruleMap,
			update:  updates[4],
			want:    false,
		},
		{
			ruleMap: ruleMap,
			update:  updates[5],
			want:    false,
		},
		{
			ruleMap: map[int][]int{
				1: {2},
				2: {3},
			},
			update: []int{3, 1, 2},
			want:   true,
		},
		{
			ruleMap: map[int][]int{
				1: {2, 3},
				2: {3},
			},
			update: []int{3, 1, 2},
			want:   false,
		},
	}
	for _, test := range tests {
		t.Run("checkUpdate", func(t *testing.T) {
			got := checkUpdate(test.update, test.ruleMap)
			if got != test.want {
				t.Errorf("%v: Expected %t, got %t", test.update, test.want, got)
			}
		})
	}
}

func TestFn2(t *testing.T) {
	t.Run("Fn2", func(t *testing.T) {
		want := 123
		got, err := fn2(testInput)
		if err != nil {
			t.Fail()
		}
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	})
}

func TestCorrectUpdate(t *testing.T) {
	section1AndSection2 := bytes.Split(testInput, []byte("\n\n"))
	section1 := bytes.Split(section1AndSection2[0], []byte("\n"))
	section2 := bytes.Split(section1AndSection2[1], []byte("\n"))
	section2 = section2[:len(section2)-1]

	ruleMap, err := prepSection1(section1)
	if err != nil {
		t.Fail()
	}

	updates, err := prepSection2(section2)
	if err != nil {
		t.Fail()
	}

	tests := []struct {
		ruleMap map[int][]int
		update  []int
		want    []int
	}{
		{
			ruleMap: ruleMap,
			update:  updates[3],
			want:    []int{97, 75, 47, 61, 53},
		},
		{
			ruleMap: ruleMap,
			update:  updates[4],
			want:    []int{61, 29, 13},
		},
		{
			ruleMap: ruleMap,
			update:  updates[5],
			want:    []int{97, 75, 47, 29, 13},
		},
	}
	for _, test := range tests {
		t.Run("checkUpdate", func(t *testing.T) {
			got := correctUpdate(test.update, test.ruleMap)
			for i, page := range got {
				if test.want[i] != page {
					t.Fatalf("%v: Expected %v, got %v", test.update, test.want, got)
				}
			}
		})
	}
}
