package lib

import "math"

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

/* --------------- find the digits of the number  --------------- */
// eg. x=18532, returns 1
func maxDigit(x int) int {
	a := 0
	for x != 0 {
		a = x
		x /= 10
	}
	return a
}

// eg. x=18532, returns 2
func leastDigit(x int) int {
	return x % 10
}
/* --------------- find the digits of the number  --------------- */
