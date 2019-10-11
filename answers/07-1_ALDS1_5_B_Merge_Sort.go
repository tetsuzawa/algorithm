package answers

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const MAX = 500

var cnt int

func round(f float64) int {
	return int(math.Floor(f + .5))
}

func Merge(A []int, n, left, mid, right int) {
	//var L = make([]int, MAX/2+2)
	//var R = make([]int, MAX/2+2)
	var L = make([]int, round(float64(n)/2))
	var R = make([]int, round(float64(n)/2))
	const sentinel = 2000000000
	n1 := mid - left
	n2 := right - mid
	for i := 0; i < n1; i++ {
		L[i] = A[left+1]
	}
	for i := 0; i < n2; i++ {
		R[i] = A[mid+1]
	}
	L = append(L, sentinel)
	R = append(R, sentinel)
	var i, j = 0, 0
	for k := left; k < right; k++ {
		cnt++
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}

func MergeSort(A []int, n, left, right int) {
	if left+1 < right {
		mid := round(float64(left + right) / 2)
		//mid := (left + right) / 2
		MergeSort(A[:], n, left, mid)
		MergeSort(A[:], n, mid, right)
		Merge(A[:], n, left, mid, right)
	}
}

func MainMergeSort() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	//A := make([]int, )
	A := make([]int, MAX)
	for i, s := range strings.Split(sc.Text(), " ") {
		A[i], _ = strconv.Atoi(s)
	}
	fmt.Println(A[:])
	MergeSort(A[:], n, 0, n)
	for i := 0; i < n; i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(A[i])
	}
	fmt.Print("\n")
	fmt.Println(cnt)
	fmt.Println(A[:])
}
