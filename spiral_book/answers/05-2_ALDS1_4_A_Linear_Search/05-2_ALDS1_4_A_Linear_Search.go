package _5_2_ALDS1_4_A_Linear_Search

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func linearSearch05_2(s []int, n int, key int) bool {
	i := 0
	s = append(s, key)
	for s[i] != key {
		i++
	}
	return i != n
}

func MainLinearSearch() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sarr := make([]int, n)
	sc.Scan()
	for i, s := range strings.Split(sc.Text(), " ") {
		sarr[i], _ = strconv.Atoi(s)
	}

	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	tarr := make([]int, q)

	sum := 0
	sc.Scan()
	for i, s := range strings.Split(sc.Text(), " ") {
		tarr[i], _ = strconv.Atoi(s)
		if linearSearch05_2(sarr, n, tarr[i]) {
			sum++
		}
	}
	fmt.Println(sum)
}
