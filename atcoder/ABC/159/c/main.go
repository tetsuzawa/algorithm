package main

import "fmt"

func main() {
	var l float64
	fmt.Scan(&l)
	fmt.Printf("%f\n", l / 3 * l / 3 * l / 3)
}
