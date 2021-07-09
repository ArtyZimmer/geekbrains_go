package fibRec

func FibRec(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	} else {
		return FibRec(number-1) + FibRec(number-2)
	}
}
