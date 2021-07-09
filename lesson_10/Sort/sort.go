package main

import (
	"fmt"
	generateRandomSlice2 "github.com/ArtyZimmer/geekbrains_go/lesson_10/Sort/generateRandomSlice"
	insertionSort2 "github.com/ArtyZimmer/geekbrains_go/lesson_10/Sort/insertionSort"
	integratedSort2 "github.com/ArtyZimmer/geekbrains_go/lesson_10/Sort/integratedSort"
)

func main() {
	var slice1, slice2 []int

	randomSlice := generateRandomSlice2.GenerateRandomSlice()
	fmt.Println("Original slice:", randomSlice)

	slice1, slice2 = randomSlice, randomSlice

	insertionSort2.InsertionSort(slice1)
	integratedSort2.IntegratedSort(slice2)

	fmt.Println("Insertion sort:", slice1)
	fmt.Println("Integrated sort:", slice2)
}
