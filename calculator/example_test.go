package main

import "fmt"

func Exampleborderlen() {
	fmt.Printf("%.2f", borderlen(30))
	// Output: 19.42
}

func Examplediameter() {
	fmt.Printf("%.2f", diameter(10))
	// Output: 3.57
}

func Exampledidgits() {
	didgits(456)
	// Output: В этом числе:\n сотен: 4,\n десятков: 5,\n единиц: 6. \n
}
