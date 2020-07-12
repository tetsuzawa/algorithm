package main

import "fmt"

var inf = 1 << 60

func main() {
	var n int
	fmt.Scan(&n)
	h := make([][]int, n)
	dp := make([][]int, n+1)
	for i := range h {
		h[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Scan(&h[i][j])
		}
	}
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				if dp[i][j]+h[i][k] > dp[i+1][k] {
					dp[i+1][k] = dp[i][j] + h[i][k]
				}
			}
		}
	}
	var res int
	for j := 0; j < 3; j++ {
		if dp[n][j] > res {
			res = dp[n][j]
		}
	}
	fmt.Println(res)
}
