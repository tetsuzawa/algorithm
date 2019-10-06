package answers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func binarySearch05_3(n, key int) bool {
	left, right := 0, n
	var mid int
	for left < right {
		mid = (left + right) / 2
		if key == sarr[mid] {
			return true
		}
		if key > sarr[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return false
}

var sarr = make([]int, 1000000)

func MainBinarySearch() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	//sarr := make([]int, n)
	sc.Scan()
	for i, s := range strings.Split(sc.Text(), " ") {
		sarr[i], _ = strconv.Atoi(s)
	}

	sc.Scan()
	_, _ = strconv.Atoi(sc.Text())

	sum := 0
	sc.Scan()
	for _, s := range strings.Split(sc.Text(), " ") {
		k, _ := strconv.Atoi(s)
		if binarySearch05_3(n, k) {
			sum++
		}
	}
	fmt.Println(sum)
}
