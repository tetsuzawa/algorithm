package main

import (
	"bufio"
	"fmt"
	"math"
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

func init() {
	nextReader = newScanner()
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func nDigits(x int) int {
	if x == 0 {
		return 1
	}
	var cnt int
	for x != 0 {
		cnt++
		x /= 10
	}
	return cnt
}

func main() {
	N, M := nextInt(), nextInt()
	s := make([]int, M, M)
	c := make([]int, M, M)
	digits := make(map[int]int)
	for i := 1; i <= N; i++ {
		digits[i] = -1
	}
	for i := 0; i < M; i++ {
		s[i], c[i] = nextInt(), nextInt()
		if digits[s[i]] != c[i] && digits[s[i]] != -1 {
			fmt.Println(-1)
			return
		}
		digits[s[i]] = c[i]
	}
	var sum int
	for i := 1; i <= N; i++ {
		v := digits[i]
		if v == -1 {
			if i == 1 && len(digits) != 1 {
				v = 1
			} else {
				v = 0
			}
		}
		sum += v * pow(10, N-i)
	}
	if nDigits(sum) != N {
		fmt.Println(-1)
		return
	}
	if sum == -1 {
		fmt.Println(0)
	} else {
		fmt.Println(sum)
	}
}

// eg. x=18532, returns 1
func maxDigit(x int) int {
	a := 0
	var cnt int
	for x != 0 {
		cnt++
		a = x
		x /= 10
	}
	return a
}
