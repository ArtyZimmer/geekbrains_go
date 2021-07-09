package insertionSort

import "testing"

func TestSort(t *testing.T) {
	testSlice := []int{2, 8, 7, 2, 0, 8, 8, 1, 2, 5}

	got := InsertionSort(testSlice)
	want := []int{0, 1, 2, 2, 2, 5, 7, 8, 8, 8}
	if !testEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}

}

func testEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
