package main

import "fmt"

func main() {
	var height, width, area float64

	fmt.Println("Эта программа посчитает площадь прямоугольника.")
	fmt.Println("Введите длину прямоугольника:")
	fmt.Scanln(&width)

	fmt.Println("Введите высоту прямоугольника:")
	fmt.Scanln(&height)

	area = height * width

	fmt.Println("Площадь Вашего прямоугольника равна", area)
}
