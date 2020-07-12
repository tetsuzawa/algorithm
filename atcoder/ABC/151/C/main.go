package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024), int(1e11))
	sc.Split(bufio.ScanWords)
	sc.Scan()
	_, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	M, _ := strconv.Atoi(sc.Text())

	var AC, penalty int
	WA := make(map[int]int)
	var hasDone = make(map[int]bool)

	for i := 0; i < M; i++ {
		sc.Scan()
		p, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		S := sc.Text()

		if hasDone[p] {
			continue
		}
		if S == "WA" {
			WA[p]++
		} else if S == "AC" {
			AC++
			penalty += WA[p]
			hasDone[p] = true
		}
	}
	fmt.Println(AC, penalty)
}
