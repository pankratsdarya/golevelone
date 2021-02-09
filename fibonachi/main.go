package main

import (
	"fmt"
)

// fibRecur returns Fibonachi number on given position using recursion
func fibRecur(number uint64) uint64 {

	switch number {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibRecur(number-1) + fibRecur(number-2)
	}
}

// map that speeds up recursion
var fibMap = map[int]uint64{0: 1, 1: 1}

// fibWMap returns Fibonachi number on given position using recursion and map
func fibWMap(number uint64) uint64 {

	switch number {
	case 0:
		return 0
	case 1:
		return 1
	default:
		_, ok := fibMap[int(number)]
		if ok {
			return fibMap[int(number)]
		}
		fibMap[int(number)] = fibWMap(number-1) + fibWMap(number-2)
	}
	return fibMap[int(number)]
}

func main() {
	var num uint64

	fmt.Print("Программа вычисляет заданное из стандартного ввода число Фибоначчи.\nСчитается, что F0=0 и F1=1. \nВведите число N: ")

	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Printf("Ошибка ввода: %v\nВам не удалось ввести число без ошибок. Программа будет закрыта.", err)
	} else {
		fmt.Printf("Вычислено рекурсией: %d\n", fibRecur(num))
		fmt.Printf("Вычислено c помощью мапы: %d", fibWMap(num))

	}
}
