package lib

import "fmt"

var ss []int
var a []int
var n int

//Recursive function for enumerating combinations
func makeCombination() {
	for i := 0; i < n; i++ {
		ss[i] = 0 // do not select i
	}
	rec(0)
}

func rec(i int) {
	if i == n {
		fmt.Println(ss)
		return
	}
	rec(i + 1)
	ss[i] = 1 // select i
	rec(i + 1)
	ss[i] = 0 // do not select i

}

//O(2^n)
//the answer will be a bit array. this is why calc cost is O(2^n)
//Recursive function for judging it can create integer or not
func solve(i, m int) bool {
	if m == 0 {
		return true
	}
	if i >= n {
		return false
	}
	res := solve(i+1, m) || solve(i+1, m-a[i])
	return res
}
