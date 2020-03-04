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

func nextInt() int { n, _ := strconv.Atoi(nextReader()); return n }

func nextInts(size int) []int {
	ns := make([]int, size)
	for i := 0; i < size; i++ {
		ns[i] = nextInt()
	}
	return ns
}

func fh(X []int, P int) int {
	var sum int
	for _, x := range X {
		sum += (x - P) * (x - P)
	}
	return sum
}

func main() {
	nextReader = newScanner()
	N := nextInt()
	X := nextInts(N)
	sort.Ints(X)
	l, r := 0, X[N-1]
	for l <= r {
		mid := l + (r-l)/2
		sl, sm, sr := fh(X, mid-1), fh(X, mid), fh(X, mid+1)
		if sl >= sm && sm <= sr {
			fmt.Println(sm)
			break
		} else if sl < sm {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
}

//for i := 1; i < X[N-1]-1; i++ {
//	sl, sm, sr := fh(X, i-1), fh(X, i), fh(X, i+1)
//	if sl < sm {
//		fmt.Println(sl)
//		break
//	} else if sm > sr {
//		fmt.Println(sm)
//		break
//	}
//}
