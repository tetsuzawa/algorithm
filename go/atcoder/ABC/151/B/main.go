package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var nextReader func() string

func newScanner() func() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024), int(1e11))
	sc.Split(bufio.ScanWords)
	return func() string {
		sc.Scan()
		return sc.Text()
	}
}

func nextInt() int { n, _ := strconv.Atoi(nextReader()); return n }

func nextInts(size int) []int {
	ns := make([]int, size)
	for i := 0; i < size; i++ {
		ns[i] = nextInt()
	}
	return ns
}

func init() {
	nextReader = newScanner()
}

func sum(xs []int) int {
	var s int
	for _, x := range xs {
		s += x
	}
	return s
}

func main() {
	N, K, M := nextInt(), nextInt(), nextInt()
	A := nextInts(N - 1)

	x := N*M - sum(A)
	if x < 0 {
		fmt.Println(0)
	} else if x > K {
		fmt.Println(-1)
	} else {
		fmt.Println(x)
	}
}
