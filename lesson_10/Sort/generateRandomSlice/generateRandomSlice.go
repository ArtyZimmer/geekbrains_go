package generateRandomSlice

import (
	"math/rand"
	"time"
)

func GenerateRandomSlice() []int {
	var randomSlice []int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randomSlice = append(randomSlice, rand.Intn(9))
	}
	return randomSlice
}
