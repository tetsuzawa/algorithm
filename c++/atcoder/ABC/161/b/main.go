package main

import (
	"fmt"
	"sort"
)

func sum(xs []int) int {
	var s int
	for _, x := range xs {
		s += x
	}
	return s
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	var a = make([]int, n)
	var t int
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		a[i] = t
	}
	s := sum(a)
	sort.Ints(a)
	for i := 0; i < m; i++ {
		if float64(a[n-i-1]) < float64(s)/float64(4*m) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
