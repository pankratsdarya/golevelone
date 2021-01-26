package main

import (
	"fmt"
	"log"
)

func main() {
	var num int

	fmt.Print("Программа выводит все простые числа от 0 до N. Считается, что 0 и 1 не являются простыми. \nВведите число N: ")
	_, err := fmt.Scanln(&num)
	if err != nil {
		log.Fatalf("Ошибка ввода: %v", err)
	}
	if num < 0 {
		log.Fatalf("Ошибка! Неверно указан диапазон!")
	}
	if num < 2 {
		log.Fatalf("Ошибка! Слишком маленький диапазон!")
	}

	arr := make([]int8, num+1)

	for i := 2; i <= num; i++ {
		for j := 2; j*i <= num; j++ {
			arr[j*i] = 1
		}
	}

	fmt.Print("Простые числа: ")

	for i := 2; i <= num; i++ {
		if arr[i] == 0 {
			fmt.Printf("%d, ", i)
		}
	}

}
