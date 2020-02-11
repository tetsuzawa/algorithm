package _3_3_ALDS1_2_A_Bubble_Sort

import (
	"bufio"
	"fmt"
	"github.com/tetsuzawa/algoLearn/answers/02-5_ALDS1_1_D_n_squ_n"
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

	_2_5_ALDS1_1_D_n_squ_n.sc = bufio.NewScanner(os.Stdin)
	_2_5_ALDS1_1_D_n_squ_n.sc.Scan()
	n, _ = strconv.Atoi(_2_5_ALDS1_1_D_n_squ_n.sc.Text())

	var a = make([]int, n)
	_2_5_ALDS1_1_D_n_squ_n.sc.Scan()
	for i, s := range strings.Split(_2_5_ALDS1_1_D_n_squ_n.sc.Text(), " ") {
		a[i], _ = strconv.Atoi(s)
	}

	BubbleSort(a, n)
}
