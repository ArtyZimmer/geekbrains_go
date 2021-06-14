package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomSlice() []int {
	var randomSlice []int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randomSlice = append(randomSlice, rand.Intn(9))
	}
	return randomSlice
}

func main() {
	randomSlice := generateRandomSlice()
	fmt.Println("Исходный массив: ", randomSlice)

	length := len(randomSlice)
	for i := 1; i < length; i++ {
		key := i
		for key > 0 {
			if randomSlice[key-1] > randomSlice[key] {
				randomSlice[key-1], randomSlice[key] = randomSlice[key], randomSlice[key-1]
			}
			key = key - 1
		}
	}
	fmt.Println("Отсортированный массив: ", randomSlice)
}
