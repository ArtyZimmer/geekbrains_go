package insertionSort

func InsertionSort(randomSlice []int) []int {
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
	return randomSlice
}
