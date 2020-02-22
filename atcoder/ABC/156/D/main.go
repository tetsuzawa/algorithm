package main

import (
	"fmt"
	"math"
)

func main() {
	var n, a, b int
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)

	var sum float64
	for i := 0; i < n; i++ {
		if i == a || i == b {
			continue
		}
		sum += combination(n, i)
	}

	fmt.Println(int(math.Mod(sum, 10e9+7)))
}

func combination(n, r int) float64 {
	if r*2 > n {
		r = n - r
	}
	dividend := 1.
	divisor := 1.
	for i := 1; i <= r; i++ {
		dividend *= float64(n) - float64(i) + 1
		divisor *= float64(i)
	}
	return dividend / divisor
}
