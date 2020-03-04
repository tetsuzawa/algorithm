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

/*
func main() {
	nextReader = newScanner()
	N := nextInt()
	var cnt int
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			is := strconv.Itoa(i)
			js := strconv.Itoa(j)
			if is[0] == js[len(js)-1] && js[0] == is[len(is)-1] {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
*/

type pair struct {
	first  int
	second int
}

func f(x int) pair {
	a := x % 10
	b := 0
	for x != 0 {
		b = x
		x /= 10
	}
	return pair{a, b}
}

func main() {
	nextReader = newScanner()
	N := nextInt()
	var freq = make(map[pair]int)
	for i := 1; i <= N; i++ {
		p := f(i)
		freq[p]++
	}
	var cnt int
	for i := 1; i <= N; i++ {
		p := f(i)
		q := pair{p.second, p.first}
		cnt += freq[q]
	}
	fmt.Println(cnt)
}
