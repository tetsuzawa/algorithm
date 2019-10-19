package _3_4_ALDS1_2_B_Selection_Sort

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//O(N^2)
func SelectionSort(a []int, n int) {
	var minj int //index of min value
	var count int

	for i := 0; i < n; i++ {
		minj = i
		for j := i; j < n; j++ {
			if a[j] < a[minj] {
				minj = j
			}
		}
		if i == minj {
			continue
		} else {
			a[i], a[minj] = a[minj], a[i]
			count++
		}
	}
	for k, p := range a {
		fmt.Print(p)
		if k == len(a)-1 {
			fmt.Printf("\n")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println(count)

}

func MainSelectionSort() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()

	var sl = make([]int, n)
	for i, s := range strings.Split(sc.Text(), " ") {
		//sl[i], _ = strconv.Atoi(s)
		sl[i], _ = strconv.Atoi(s)
	}

	SelectionSort(sl, n)
}
