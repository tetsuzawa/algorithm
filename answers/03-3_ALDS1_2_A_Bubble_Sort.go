package answers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(a []int, n int) {
	var isArranged = true
	var j int

	for isArranged {
		isArranged = false
		for j = n - 1; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				isArranged = true
			}
		}
		fmt.Println(a)
	}
}

func MainBubbleSort() {
	var n int

	sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())

	var a = make([]int, n)
	sc.Scan()
	for i, s := range strings.Split(sc.Text(), " ") {
		a[i], _ = strconv.Atoi(s)
	}

	BubbleSort(a, n)
}
