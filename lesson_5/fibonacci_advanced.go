package main

import (
	"fmt"
	"os"
)

func fibonacci(input int) int {
	if input < 2 {
		return 1
	}
	return fibonacci(input-2) + fibonacci(input-1)
}

func calculations(input int) int {
	var result int
	result = fibonacci(input)
	return result
}

func main() {
	var input, result, answer int
	runStatus := true
	data := make(map[int]int)
	fmt.Println("Введите номер числа Фибоначчи:")
	fmt.Scanln(&input)
	for runStatus {
		if _, ok := data[input]; ok {
			fmt.Println(input, "число Фибоначчи —", data[input])
			fmt.Println("Поищем другое число? (0 — да, 1 — нет)")
			fmt.Scan(&answer)
			if answer == 0 {
				fmt.Println("Введите номер числа Фибоначчи:")
				fmt.Scanln(&input)
				continue
			} else {
				os.Exit(1)
			}
		} else {
			result = calculations(input)
			data[input] = result
			fmt.Println(input, "число Фибоначчи —", result)
			fmt.Println("Поищем другое число? (0 — да, 1 — нет)")
			fmt.Scan(&answer)
			if answer == 0 {
				fmt.Println("Введите номер числа Фибоначчи:")
				fmt.Scanln(&input)
				continue
			} else {
				os.Exit(1)
			}
		}
	}
}
