package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// theory, faster
func main() {
	nextReader = newScanner()
	is := nextInts(3)
	ni, a, b := is[0], is[1], is[2]
	var total int
	for i := 1; i <= ni; i++ {
		sum := findSumOfDigits(i)
		if a <= sum && sum <= b {
			total += i
		}
	}
	fmt.Println(total)
}

func findSumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

//func main() {
//	nextReader = newScanner()
//	is := nextInts(3)
//	ni, a, b := is[0], is[1], is[2]
//	var sum int
//	for i := 1; i <= ni; i++ {
//		dsum := 0
//		for _, v := range strconv.Itoa(i) {
//			digit, _ := strconv.Atoi(string(v))
//			dsum += digit
//		}
//		if a <= dsum && dsum <= b {
//			sum += i
//		}
//	}
//	fmt.Println(sum)
//}

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
