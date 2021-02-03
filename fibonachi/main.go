package main

import (
	"fmt"
	"log"
)

// fibRecur возвращает число Фибоначчи, полученное с помощью рекурсии
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

// мапа для ускорения рекурсии
var fibMap = map[int]uint64{0: 1, 1: 1}

// fibWMap возвращает число Фибоначчи, полученное с помощью рекурсии, ускоренной мапой
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
		log.Fatalf("Ошибка ввода: %v", err)
	}

	fmt.Printf("Вычислено рекурсией: %d\n", fibRecur(num))
	fmt.Printf("Вычислено c помощью мапы: %d", fibWMap(num))

}
