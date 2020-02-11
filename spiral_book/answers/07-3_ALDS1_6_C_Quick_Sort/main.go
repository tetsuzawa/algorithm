package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX      = 100000
	SENTINEL = 20000000000
)

type Card struct {
	Suit  string
	Value int
}

var L = make([]Card, MAX/2+2)
var R = make([]Card, MAX/2+2)

func merge(a []Card, n, left, mid, right int) {
	n1 := mid - left
	n2 := right - mid
	for i := 0; i < n1; i++ {
		L[i] = a[left+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = a[mid+i]
	}
	i, j := 0, 0
	L[n1].Value = SENTINEL
	R[n2].Value = SENTINEL
	for k := left; k < right; k++ {
		if L[i].Value <= R[j].Value {
			a[k] = L[i]
			i++
		} else {
			a[k] = R[j]
			j++
		}
	}
}

func mergeSort(a []Card, n, left, right int) {
	var mid int
	if left+1 < right {
		mid = (left + right) / 2
		mergeSort(a, n, left, mid)
		mergeSort(a, n, mid, right)
		merge(a, n, left, mid, right)
	}

}

func partition(a []Card, n, p, r int) int {
	x := a[r]
	i := p - 1

	for j := p; j < r; j++ {
		if a[j].Value <= x.Value {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func quickSort(a []Card, n, p, r int) {
	if p < r {
		q := partition(a, n, p, r)
		quickSort(a, n, p, q-1)
		quickSort(a, n, q+1, r)
	}
}

func main() {
	a := make([]Card, MAX)
	b := make([]Card, MAX)
	isStable := true

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for i := 0; i < n; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), " ")
		a[i].Suit = s[0]
		b[i].Suit = s[0]
		a[i].Value, _ = strconv.Atoi(s[1])
		b[i].Value, _ = strconv.Atoi(s[1])
	}

	mergeSort(a, n, 0, n)
	quickSort(b, n, 0, n-1)

	for i := 0; i < n; i++ {
		if a[i].Suit != b[i].Suit {
			isStable = false
		}
	}

	if isStable {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
	for i := 0; i < n; i++ {
		fmt.Println(b[i].Suit, b[i].Value)
	}
}
