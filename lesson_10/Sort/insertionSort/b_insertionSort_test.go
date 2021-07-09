package insertionSort

import (
	generateRandomSlice2 "github.com/ArtyZimmer/geekbrains_go/lesson_10/Sort/generateRandomSlice"
	integratedSort2 "github.com/ArtyZimmer/geekbrains_go/lesson_10/Sort/integratedSort"
	"testing"
)

var sink []int

func BenchmarkInsertionSort(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = InsertionSort(generateRandomSlice2.GenerateRandomSlice())
	}
	sink = res
}

func BenchmarkIntegratedSort(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = integratedSort2.IntegratedSort(generateRandomSlice2.GenerateRandomSlice())
	}
	sink = res
}
