package _6_2_ALDS1_5_A_Exhastive_Search

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n int
var a = make([]int, 50)

func solve(i, m int) int {
	if m == 0 {
		return 1
	}
	if i >= n {
		return 0
	}
	res := solve(i+1, m)
	if r := solve(i+1, m-a[i]); r != 0 {
		res = r
	}
	return res
}

func MainExhaustiveSearch() {
	var M int

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	for i, s := range strings.Split(sc.Text(), " ") {
		a[i], _ = strconv.Atoi(s)
	}
	sc.Scan()
	_, _ = strconv.Atoi(sc.Text())
	for _, s := range strings.Split(sc.Text(), " "){
		M, _ = strconv.Atoi(s)
		if solve(0, M) == 1{
			fmt.Println("yes")
		}else {
			fmt.Println("no")
		}
	}
}
