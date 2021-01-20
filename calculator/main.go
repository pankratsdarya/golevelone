package main

import (
	"fmt"
	"math"
)

// borderlen возвращает длину окружности
func borderlen(s float64) float64 {
	return math.Sqrt(s/math.Pi) * 2 * math.Pi
}

// diameter возвращает диаметр окружности
func diameter(s float64) float64 {
	return 2 * math.Sqrt(s/math.Pi)
}

// didgits разбирает число на сотни, десятки и единицы
func didgits(num uint32) {
	var hundreds, tens, units uint8
	hundreds = uint8(math.Floor(float64(num) / 100))
	num -= uint32(hundreds) * 100
	tens = uint8(math.Floor(float64(num) / 10))
	units = uint8(num % 10)
	fmt.Printf("В этом числе:\n сотен: %d,\n десятков: %d,\n единиц: %d. \n", hundreds, tens, units)
}

func main() {
	//Расчет площади прямоугольника
	var height, width int32
	fmt.Println("Расчет площади прямоугольника.")
	fmt.Print("Введите высоту: ")
	fmt.Scanln(&height)
	fmt.Print("Введите ширину: ")
	fmt.Scanln(&width)
	fmt.Printf("Площадь прямоугольника равна %d. \n", height*width)
	//Расчет диаметра и длины окружности по заданной площади круга
	var square float64
	fmt.Println("Расчет диаметра и длины окружности по заданной площади круга.")
	fmt.Print("Введите площадь: ")
	fmt.Scanln(&square)
	fmt.Printf("Диаметр круга равен %.4f. \n", diameter(square))
	fmt.Printf("Длина окружности равна %.4f. \n", borderlen(square))
	//Разбор числа на сотни, десятки и единицы
	var newnumber uint32
	fmt.Println("Разбор числа на сотни, десятки и единицы.")
	fmt.Print("Введите число от 100 до 999: ")
	fmt.Scanln(&newnumber)
	didgits(newnumber)
}
