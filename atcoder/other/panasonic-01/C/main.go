package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	// if (c-a-b)*(c-a-b) > 4*a*b {
	if (c - a - b) > 2*int(math.Sqrt(float64(a*b))) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
