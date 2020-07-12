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

	N := nextInt()
	ai := nextInts(N)

	num := 0
	for _, v := range ai {
		if v == num+1 {
			num++
		}
	}
	cnt := -1
	if num != 0 {
		cnt = N - num
	}
	fmt.Println(cnt)
}

/*
func main() {
	nextReader = newScanner()

	N := nextInt()
	ai := nextInts(N)
	bucket := make(map[int]int)
	for i := 1; i <= N; i++ {
		bucket[i] = -1
	}
	for i, v := range ai {
		bucket[v] = i
	}
	cnt := -1
	var idx int
	for i := 1; i <= N; i++ {
		if bucket[i] == -1 {
			if i == 1 {
				fmt.Println(-1)
				break
			} else {
				fmt.Println(i - 1)
			}
			//TODO (break)
		}
		if bucket[i] > idx {
			idx = bucket[i]
			cnt += 1
		}
	}
}

*/
