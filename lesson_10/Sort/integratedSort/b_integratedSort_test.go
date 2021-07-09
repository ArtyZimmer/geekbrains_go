package integratedSort

import (
	generateRandomSlice2 "github.com/ArtyZimmer/geekbrains_go/lesson_10/Sort/generateRandomSlice"
	"testing"
)

const N = 20

var sink []int

func BenchmarkIntegratedSort(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = IntegratedSort(generateRandomSlice2.GenerateRandomSlice())
	}
	sink = res
}
