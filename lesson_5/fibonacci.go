package main

import "fmt"

func fibonacci(input int) int {
	if input < 2 {
		return 1
	}
	return fibonacci(input-2) + fibonacci(input-1)
}

func main() {
	var input int
	fmt.Print("Введите номер числа последовательности Фибоначчи: ")
	fmt.Scan(&input)
	fmt.Println(input, "число Фибоначчи —", fibonacci(input))
}
