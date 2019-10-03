package answers

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func MainLinearSearch() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sarr := make([]int, n)
	for i, s := range strings.Split(sc.Text(), " ") {
		sarr[i], _ = strconv.Atoi(s)
	}
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	tarr := make([]int, q)
	for i, s := range strings.Split(sc.Text(), " ") {
		tarr[i], _ = strconv.Atoi(s)
	}



}
