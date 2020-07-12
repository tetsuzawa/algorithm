package main

import "fmt"

var inf = 1 << 60

func main() {
	var n, w int
	fmt.Scan(&n)
	fmt.Scan(&w)
	ws := make([]int, n)
	vs := make([]int, n)
	dp := make([][]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&ws[i])
		fmt.Scan(&vs[i])
	}
	for i := range dp {
		dp[i] = make([]int, 100100)
	}

	for i := 0; i < n; i++ {
		for sumW := 0; sumW <= w; sumW++ {
			// i 番目の品物を選ぶ場合
			if sumW-ws[i] >= 0 {
				if dp[i][sumW-ws[i]]+vs[i] > dp[i+1][sumW] {
					dp[i+1][sumW] = dp[i][sumW-ws[i]] + vs[i]
				}
			}
			// i 番目の品物を選ばない場合
			if dp[i][sumW] > dp[i+1][sumW] {
				dp[i+1][sumW] = dp[i][sumW]
			}
		}
	}
	fmt.Println(dp[n][w])
}
