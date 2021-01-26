package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var a, b, res float32
	var op, zero string

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Fatalf("Ошибка ввода: %v", err)
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		log.Fatalf("Ошибка ввода: %v", err)
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /): ")
	_, err = fmt.Scanln(&op)
	if err != nil {
		log.Fatalf("Ошибка ввода: %v", err)
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Print("Внимание! Деление на 0! Продолжить? y/n  ")
			fmt.Scanln(&zero)
			if zero == "y" {
				res = a / b
			} else if zero == "n" {
				fmt.Println("Операция деления отменена")
				os.Exit(1)
			} else {
				log.Fatal("Операция выбрана неверно")
			}

		} else {
			res = a / b
		}
	default:
		log.Fatal("Операция выбрана неверно")
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}
