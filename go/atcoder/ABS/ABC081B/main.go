package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	nextReader = newScanner()
	n := nextInt()
	is := nextInts(n)
	var cnt int
	func() {
		for {
			for i, _ := range is {
				if isOdd(is[i]) {
					return
				}
				is[i] /= 2
			}
			cnt++
		}
	}()
	// used flag
	/*
		hasFoundOdd := false
		for {
			for i, _ := range is {
				if isOdd(is[i]) {
					hasFoundOdd = true
				}
				is[i] /= 2
			}
			if hasFoundOdd {
				break
			}
			cnt++
		}
	*/
	fmt.Println(cnt)
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

func isOdd(x int) bool {
	return x%2 != 0
}
