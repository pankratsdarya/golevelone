package main

import "testing"

func TestFibRecur(t *testing.T) {
	tests := []struct {
		name  string
		input uint64
		want  uint64
	}{
		{name: "zero", input: 0, want: 0},
		{name: "one", input: 1, want: 1},
		{name: "five", input: 5, want: 5},
		{name: "ten", input: 10, want: 55},
	}

	for _, tc := range tests {
		got := fibRecur(tc.input)
		if tc.want != got {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func TestFibWMap(t *testing.T) {
	tests := []struct {
		name  string
		input uint64
		want  uint64
	}{
		{name: "zero", input: 0, want: 0},
		{name: "one", input: 1, want: 1},
		{name: "five", input: 5, want: 5},
		{name: "ten", input: 10, want: 55},
	}

	for _, tc := range tests {
		got := fibWMap(tc.input)
		if tc.want != got {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
