package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func nextString() string { return nextReader() }

func nextInt() int { n, _ := strconv.Atoi(nextReader()); return n }

func nextInts(size int) []int {
	ns := make([]int, size)
	for i := 0; i < size; i++ {
		ns[i] = nextInt()
	}
	return ns
}

func main() {
	nextReader = newScanner()
	N, K := nextInt(), nextInt()
	Ai := nextInts(N)
	sort.Ints(Ai)

	var i = 0
	var l = 0
	var r = N - 1
	var ans = Ai[l] * Ai[r]
	//var cnt = 0
	for i = 0; i < K; i++ {
		if Ai[l]*Ai[r-1] < Ai[l+1]*Ai[r] {
			r -= 1
			ans = Ai[l] * Ai[r-1]
			//K = K - (N - 1)
			//N -= 1
		} else {
			l += 1
			ans = Ai[l+1] * Ai[r]
			//K = K - (N - 1)
			//N -= 1
		}
		i++
	}
	fmt.Println(ans)
}
