package main

import "testing"

var result []int

func BenchmarkBubble(b *testing.B) {
	input := []int{7, 2, 0, 4, 9, 8, 6, 1, 3, 5}
	res := make([]int, 10)

	for i := 0; i < b.N; i++ {
		res = bubble(input)
	}

	copy(result, res)
}

func BenchmarkInsertion(b *testing.B) {
	input := []int{7, 2, 0, 4, 9, 8, 6, 1, 3, 5}
	res := make([]int, 10)

	for i := 0; i < b.N; i++ {
		res = insertion(input)
	}

	copy(result, res)
}
