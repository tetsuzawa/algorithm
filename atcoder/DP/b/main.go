package main

import "fmt"

var inf = 1 << 60

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	h := make([]int, n)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
	}
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := 1; j <= k; j++ {
			if i-j < 0 {
				break
			}
			if dp[i-j]+abs(h[i]-h[i-j]) < dp[i] {
				dp[i] = dp[i-j] + abs(h[i]-h[i-j])
			}
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
