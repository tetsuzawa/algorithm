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

type set struct {
	c int
	s string
}

type sets []set

func (ss sets) Len() int {
	return len(ss)
}

func (ss sets) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss sets) Less(i, j int) bool {
	if ss[i].c == ss[j].c {
		return ss[i].s < ss[j].s
	} else {
		return ss[i].c > ss[j].c
	}
}

func main() {
	nextReader = newScanner()
	N := nextInt()
	var S = make(map[string]int, N)
	for i := 0; i < N; i++ {
		S[nextString()] += 1
	}

	ss := sets{}
	for k, v := range S {
		e := set{v, k}
		ss = append(ss, e)
	}
	sort.Sort(ss)
	for i := 0; ss[i].c == ss[0].c; i++ {
		fmt.Println(ss[i].s)
		if i == len(ss)-1 {
			break
		}
	}
}
