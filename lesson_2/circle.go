package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter, length, area float64

	fmt.Println("Эта программа вычисляет диаметр и длину окружности по её площади. " +
		"Введите площадь окружности:")
	fmt.Scanln(&area)

	diameter = (math.Sqrt(area / math.Pi)) * 2
	length = math.Pi * diameter

	fmt.Println("Вы задали площадь окружности равную:", area)
	fmt.Print("Диаметр такой окружности равен ", fmt.Sprintf("%.2f", diameter), ", "+
		"а длина 20", fmt.Sprintf("%.2f", length))
}
