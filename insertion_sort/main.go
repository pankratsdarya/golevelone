package main

import "fmt"

// bubble sort
func bubble(src []int) []int {
	dest := make([]int, len(src))
	copy(dest, src)
	for j := 1; j < len(src)-1; j++ {
		for i := 0; i < len(src)-j; i++ {
			if dest[i] > dest[i+1] {
				dest[i], dest[i+1] = dest[i+1], dest[i]
			}
		}
	}
	return dest
}

// insertions sort
func insertion(src []int) []int {
	dest := make([]int, len(src))
	copy(dest, src)
	for i := 2; i < len(dest); i++ {
		for j := i; j > 0 && dest[j] < dest[j-1]; j-- {
			dest[j], dest[j-1] = dest[j-1], dest[j]
		}
	}
	return dest
}

func main() {

	arrOne := []int{7, 2, 0, 4, 9, 8, 6, 1, 3, 5}
	arrTwo := []int{0, 4, 7, 2, 6, 1, 3, 5, 9, 8}

	arrOneSort := bubble(arrOne)
	arrTwoSort := insertion(arrTwo)

	fmt.Println(arrOne)
	fmt.Println(arrOneSort)
	fmt.Println(arrTwo)
	fmt.Println(arrTwoSort)
}
