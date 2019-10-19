package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var a = make([]int, 100000)

func Partition(a []int, p, r int) int {
	x := a[r]
	i := p - 1

	for j := p; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	var n int
	n, _ = strconv.Atoi(sc.Text())

	sc.Scan()
	for i, s := range strings.Split(sc.Text(), " ") {
		a[i], _ = strconv.Atoi(s)
	}
	q := Partition(a, 0, n-1)

	for i := 0; i < n; i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		if i == q {
			fmt.Print("[")
		}
		fmt.Print(a[i])
		if i == q {
			fmt.Print("]")
		}
	}
	fmt.Printf("\n")
}
