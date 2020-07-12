package main

import "fmt"

var INF = 1 << 60

var dp = make([]int, 100010)

func main() {
	var n int
	fmt.Scan(&n)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
	}
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0

	for i := 1; i < n; i++ {
		if dp[i] > dp[i-1]+abs(h[i]-h[i-1]) {
			dp[i] = dp[i-1] + abs(h[i]-h[i-1])
		}
		if i < 2 {
			continue
		}
		if dp[i] > dp[i-2]+abs(h[i]-h[i-2]) {
			dp[i] = dp[i-2] + abs(h[i]-h[i-2])
		}
	}
	fmt.Println(dp[n-1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	};
	return a
}
