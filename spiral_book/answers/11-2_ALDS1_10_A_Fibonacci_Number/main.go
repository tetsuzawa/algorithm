package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var dp = make([]int, 50)

func fibonacci_slow(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibonacci_slow(n-2) + fibonacci_slow(n-1)
}

func fib(n int) int {
	if n == 0 || n == 1 {
		dp[n] = 1
		return dp[n]
	}
	if dp[n] != -1 {
		return dp[n]
	}
	dp[n] = fib(n-1) + fib(n-2)
	return dp[n]
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < 50; i++ {
		dp[i] = -1
	}
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	fmt.Println(fib(n))
}
