package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var checkVarType string
	var checkVarTypeBool bool

	fmt.Println("Эта программа найдёт все простые числа в диапазоне от 0 до N")
	fmt.Println("Введите число N:")
	fmt.Scanln(&checkVarType)
	checkVarTypeBool = true
	for checkVarTypeBool {
		if s, err := strconv.Atoi(checkVarType); err == nil {
			n = s
			checkVarTypeBool = false
		} else {
			fmt.Println("Вы ввели недопустимое значение! Допускаются только числа. Повторите ввод:")
			fmt.Scanln(&checkVarType)
		}
	}
	var slice []int
	for i := 0; i <= n; i++ {
		slice = append(slice, i)
	}

}
