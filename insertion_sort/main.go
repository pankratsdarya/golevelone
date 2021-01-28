package main

import "fmt"

func main() {

	arr := [...]int{0, 4, 7, 2, 6, 1, 3, 5, 9, 8}

	for i := 2; i < cap(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	fmt.Println(arr)
}
