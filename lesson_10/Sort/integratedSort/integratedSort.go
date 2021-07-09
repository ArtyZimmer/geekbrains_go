package integratedSort

import "sort"

func IntegratedSort(randomSlice []int) []int {
	sort.Ints(randomSlice)
	return randomSlice
}
