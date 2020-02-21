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

//func pow(x, y int) int {
//	return int(math.Pow(float64(x), float64(y)))
//}

func f(H int) int {
	if H > 1 {
		return 2*f(H/2) + 1
	}
	return 1
}

func main() {
	nextReader = newScanner()
	H := nextInt()
	fmt.Println(f(H))
}

//func main() {
//	nextReader = newScanner()
//	H := nextInt()
//
//	cnt := 1
//	for H/2 >= 1 {
//		H /= 2
//		cnt++
//	}
//	var sum int
//	for i := 0; i < cnt; i++ {
//		sum += pow(2, i)
//	}
//	fmt.Println(sum)
//}
