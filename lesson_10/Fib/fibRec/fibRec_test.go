package fibRec

import "testing"

func TestFibRec(t *testing.T) {
	table := []struct {
		arg  int
		want int
	}{
		{0, 0},
		{10, 55},
		{20, 6765},
		{30, 832040},
	}
	for _, entry := range table {
		got := FibRec(entry.arg)
		if got != entry.want {
			t.Errorf("For %d got %d, want %d", entry.arg, got, entry.want)
		}
	}
}
