package main

import (
	"bytes"
	"testing"
)

var testInput1 = []byte(`mul(4*
mul(6,9!
?(12,34)
xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
`)

var testInput2 = []byte(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)
+mul(32,64](mul(11,8)undo()?mul(8,5))
`)

func TestFn1(t *testing.T) {
	t.Run("Fn1", func(t *testing.T) {
		want := 161
		got, err := fn1(testInput1)
		if err != nil {
			t.Fail()
		}
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	})
}

func TestScanLine(t *testing.T) {
	testInputLines := bytes.Split(testInput1, []byte("\n"))
	testLines := testInputLines[:len(testInputLines)-1]
	tests := []struct {
		line []byte
		want int
	}{
		{
			line: testLines[0],
			want: 0,
		},
		{
			line: testLines[1],
			want: 0,
		},
		{
			line: testLines[2],
			want: 0,
		},
		{
			line: testLines[3],
			want: 161,
		},
	}
	for _, test := range tests {
		t.Run("scanLine1", func(t *testing.T) {
			got, err := scanLine1(test.line)
			if err != nil {
				t.Fail()
			}
			if got != test.want {
				t.Errorf("%s: Expected %d, got %d", test.line, test.want, got)
			}
		})
	}
}

func TestFn2(t *testing.T) {
	t.Run("Fn2", func(t *testing.T) {
		want := 48
		got, err := fn2(testInput2)
		if err != nil {
			t.Fail()
		}
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	})
}

func TestScanLine2(t *testing.T) {
	testInputLines := bytes.Split(testInput2, []byte("\n"))
	testLines := testInputLines[:len(testInputLines)-1]
	tests := []struct {
		line        []byte
		enabled     bool
		want        int
		wantEnabled bool
	}{
		{
			line:        testLines[0],
			enabled:     true,
			want:        8,
			wantEnabled: false,
		},
		{
			line:        testLines[0],
			enabled:     false,
			want:        0,
			wantEnabled: false,
		},
		{
			line:        testLines[1],
			enabled:     true,
			want:        128,
			wantEnabled: true,
		},
		{
			line:        testLines[1],
			enabled:     false,
			want:        40,
			wantEnabled: true,
		},
		{
			line:        []byte("mul(1,10)don't()do()mul(1,3)"),
			enabled:     true,
			want:        13,
			wantEnabled: true,
		},
	}
	for _, test := range tests {
		t.Run("scanLine1", func(t *testing.T) {
			got, enabled, err := scanLine2(test.line, test.enabled)
			if err != nil {
				t.Fail()
			}
			if enabled != test.wantEnabled {
				t.Errorf("%s: Expected %t, got %t", test.line, test.wantEnabled, enabled)
			}
			if got != test.want {
				t.Errorf("%s: Expected %d, got %d", test.line, test.want, got)
			}
		})
	}
}
