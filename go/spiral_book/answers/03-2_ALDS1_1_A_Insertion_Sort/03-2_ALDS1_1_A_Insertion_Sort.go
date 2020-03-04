package _3_2_ALDS1_1_A_Insertion_Sort

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertionSort(a []int, n int) {
	var i, j int
	var v int

	for k, p := range a {
		fmt.Printf("%d", p)
		if k == len(a)-1 {
			fmt.Printf("\n")
		} else {
			fmt.Print(" ")
		}
	}

	for i = 1; i < n; i++ {
		v = a[i]
		j = i - 1
		for j >= 0 && a[j] > v {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = v
		for k, p := range a {
			fmt.Printf("%d", p)
			if k == len(a)-1 {
				fmt.Printf("\n")
			} else {
				fmt.Print(" ")
			}
		}
	}
}

func MainInsertionSort() {
	var n int
	var s string

	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	s = sc.Text()

	var a = make([]int, n)

	for i, t := range strings.Split(s, " ") {
		a[i], _ = strconv.Atoi(t)
	}

	insertionSort(a, n)
}
