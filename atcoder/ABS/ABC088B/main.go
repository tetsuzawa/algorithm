package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	nextReader = newScanner()
	n := nextInt()
	as := nextInts(n)
	sort.Sort(sort.Reverse(sort.IntSlice(as)))
	var sum int
	// theory, readable
	for i := 0; i < n; i++ {
		if isEven(i) {
			sum += as[i]
		} else {
			sum -= as[i]
		}
	}
	/*
		// ターンは偶数奇数
			for i := 0; i < n; i += 2 {
				sum += as[i]
				if i+1 == n {
					break
				}
				sum -= as[i+1]
			}
	*/
	fmt.Println(sum)
}

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

func isEven(x int) bool {
	return x%2 == 0
}
