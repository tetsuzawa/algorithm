package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	nextReader = newScanner()
	n := nextInt()
	var di = make([]int, n)
	for i := 0; i < n; i++ {
		di[i] = nextInt()
	}
	// ************ set type, AC ************

	set := make(map[int]struct{})
	for _, v := range di {
		set[v] = struct{}{}
	}
	fmt.Println(len(set))

	// ************ set type ************
	// ************ bucket, AC ************
	/*
		var num [101]int
		for _, v := range di {
			num[v]++
		}

		steps := 0
		for i := 1; i <= 100; i++ {
			if num[i] > 0 {
				steps++
			}
		}
		fmt.Println(steps)
	*/
	// ************ bucket ************

	// ************ basic AC ************
	/*
		sort.Ints(di)
			steps := 1
			d := di[0]
			for _, v := range di {
				if d < v {
					d = v
					steps++
				}
			}
			fmt.Println(steps)

	*/
	// ************ basic AC ************
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
