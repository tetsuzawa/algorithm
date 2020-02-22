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

func main() {
	nextReader = newScanner()
	S := nextString()
	T := nextString()
	var balls = make(map[string]int, 0)
	is := nextInts(2)
	balls[S] = is[0]
	balls[T] = is[1]
	if nextString() == S {
		balls[S] -= 1
	} else {
		balls[T] -= 1
	}
	fmt.Println(balls[S], balls[T])
}
