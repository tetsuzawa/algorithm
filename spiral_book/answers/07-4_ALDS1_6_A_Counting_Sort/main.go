package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX  = 200001
	VMAX = 10000
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var c = make([]int, VMAX+1)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	var a = make([]uint16, n+1)
	var b = make([]uint16, n+1)

	sc.Scan()
	for i, s := range strings.Split(sc.Text(), " ") {
		in, _ := strconv.Atoi(s)
		a[i+1] = uint16(in)
		c[a[i]]++
	}

	for i := 1; i <= VMAX; i++ {
		c[i] = c[i] + c[i-1]
	}

	for j := 1; j <= n; j++ {
		b[c[a[j]]] = a[j]
		c[a[j]]--
	}

	for i := 1; i <= n; i++ {
		if i > 1 {
			fmt.Print(" ")
		}
		fmt.Print(b[i])
	}
	fmt.Printf("\n")
}
