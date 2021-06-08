package main

import (
	"fmt"
)

func main() {
	var number, units, tens, hundreds int

	fmt.Println("Эта программа отображает количество единиц, десятков и сотен в введённом Вами числе. " +
		"Введите целое трёхзначное число:")
	fmt.Scanln(&number)

	hundreds = number / 100
	units = number % 10
	tens = (number - (hundreds * 100) - units) / 10

	fmt.Println("Количество сотен:", hundreds)
	fmt.Println("Количество десятков:", tens)
	fmt.Println("Количество единиц:", units)
}
