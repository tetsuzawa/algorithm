package main

import "fmt"

// eg. x=18532, returns 5
func nDigits(x int) int {
	var cnt int
	for x != 0 {
		cnt++
		x /= 10
	}
	return cnt
}

func main() {
	a := 18532
	cnt := nDigits(a)
	fmt.Println(cnt)
}
