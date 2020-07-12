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

func nextString() string { return nextReader() }

func nextInt() int { n, _ := strconv.Atoi(nextReader()); return n }

func nextInts(size int) []int {
	ns := make([]int, size)
	for i := 0; i < size; i++ {
		ns[i] = nextInt()
	}
	return ns
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func init() {
	nextReader = newScanner()
}

func main() {
	A := make([][]int, 3, 3)
	for i := 0; i < 3; i++ {
		A[i] = make([]int, 3, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			A[i][j] = nextInt()
		}
	}
	N := nextInt()
	hits := make([]int, 3, 3)
	for k := 0; k < N; k++ {
		b := nextInt()
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if A[i][j] == b {
					hits = append(hits, 3*i+(j+1))
				}
			}
		}
	}
	for i := 1; i <= 3; i++ {
		if contains(hits, i) && contains(hits, i+3) && contains(hits, i+6) {
			fmt.Println("Yes")
			return
		} else if contains(hits, 3*i-2) && contains(hits, 3*i-1) && contains(hits, 3*i) {
			fmt.Println("Yes")
			return
		}
	}
	if contains(hits, 1) && contains(hits, 5) && contains(hits, 9) {
		fmt.Println("Yes")
		return
	}
	if contains(hits, 3) && contains(hits, 5) && contains(hits, 7) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")

}
