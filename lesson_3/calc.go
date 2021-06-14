package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result, number_1, number_2 float64
	var operation, checkVarType string
	var checkVerTypeBool bool
	operationsList := []string{"+", "-", "*", "/", "%", "^"}

	fmt.Println("Это программа-калькулятор!")
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&checkVarType)
	checkVerTypeBool = true
	for checkVerTypeBool {
		if s, err := strconv.Atoi(checkVarType); err == nil {
			number_1 = float64(s)
			checkVerTypeBool = false
		} else {
			fmt.Println("Введено некорректное значение переменной (ожидается число)! Повторите ввод:")
			fmt.Scanln(&checkVarType)
		}
	}
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&checkVarType)
	checkVerTypeBool = true
	for checkVerTypeBool {
		if s, err := strconv.Atoi(checkVarType); err == nil {
			number_2 = float64(s)
			checkVerTypeBool = false
		} else {
			fmt.Println("Введено некорректное значение переменной (ожидается число)! Повторите ввод:")
			fmt.Scanln(&checkVarType)
		}
	}
	fmt.Println("Теперь введите желаему операцию:", strings.Join(operationsList, ", "))
	fmt.Scanln(&operation)

	switch operation {
	case operationsList[0]:
		result = number_1 + number_2
	case operationsList[1]:
		result = number_1 - number_2
	case operationsList[2]:
		result = number_1 * number_2
	case operationsList[3]:
		for number_2 == 0 {
			fmt.Println("На 0 делить нельзя! Введите другое второе число: ")
			fmt.Scanln(&number_2)
		}
		result = number_1 / number_2
	case operationsList[4]:
		for number_2 == 0 {
			fmt.Println("На 0 делить нельзя! Введите другое второе число: ")
			fmt.Scanln(&number_2)
		}
		for number_1 >= number_2 {
			number_1 -= number_2
			result = number_1
		}
		if number_1 < number_2 {
			result = number_1
		}
	case operationsList[5]:
		result = math.Pow(number_1, number_2)
	default:
		fmt.Println("Введена некорректная операция, программа будет завершена!")
		os.Exit(1)
	}
	fmt.Println("Результат выполнения операции:", result)
}
