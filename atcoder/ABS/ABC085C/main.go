package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	nextReader = newScanner()
	is := nextInts(2)
	n, Y := is[0], is[1]

	var (
		res10000 = -1
		res5000  = -1
		res1000  = -1
	)

	func() {
		for x := 0; x <= n; x++ {
			for y := 0; y <= n-x; y++ {
				z := n - x - y
				sum := 10000*x + 5000*y + 1000*z
				if sum == Y {
					res10000 = x
					res5000 = y
					res1000 = z
					return
				}
			}
		}
	}()
	fmt.Println(res10000, res5000, res1000)

	/*
		// too slow (N+1)^3
		func() {
			for x := 0; x <= n; x++ {
				for y := 0; y <= n; y++ {
					for z := 0; z <= n; z++ {
						sum := 10000*x + 5000*y + 1000*z
						if sum == Y && x+y+z == n {
							fmt.Println(x, y, z)
							return
						}
					}
				}
			}
			fmt.Println(-1, -1, -1)
		}()
	*/
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
