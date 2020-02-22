package main

import (
	"fmt"
)

func main() {
	n := 4
	r := 3
	is := nCr(n, r)
	fmt.Printf("%v, %v: %v\n", n, r, is[n][r])
}

func nCr(n, r int) [][]int {
	is := make([][]int, n+1)
	for i, _ := range is {
		is[i] = make([]int, n+1)
	}
	for i := 0; i < len(is); i++ {
		is[i][0] = 1
		is[i][i] = 1
	}
	for j := 1; j < len(is); j++ {
		for k := 1; k < j; k++ {
			is[j][k] = is[j-1][k-1] + is[j-1][k]
		}
	}
	return is
}

func bits() {
	n := 3
	for bit := 0; bit < (1 << n); bit++ {
		S := []int{}
		for i := 0; i < n; i++ {
			if bit&(1<<i) != 0 {
				S = append(S, i)
			}
		}
		// Print
		fmt.Print(bit, ": {")
		for i := 0; i < len(S); i++ {
			fmt.Print(S[i], " ")
		}
		fmt.Println("}")
	}
}
