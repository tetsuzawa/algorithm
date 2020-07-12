package algorithm

import (
	"fmt"
	"math"
)

var ss []int
var a []int
var n int

/************ Exhaustive search *************/
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

/************ Exhaustive search *************/

/************ Koch curve *************/

type Point struct {
	x float64
	y float64
}

// d: deepness of recursion
// p1, p2: points of edges
func koch(d int, p1, p2 Point) {
	if d == 0 {
		return
	}

	var s, t Point // Points equally divided into 3
	var u Point    // Vertex of equilateral triangle

	// calc point of s, u, t by using values of p1 and p2
	th := math.Pi * 60.0 / 180.0
	s.x = (2.0*p1.x + 1.0*p2.x) / 3.0
	s.y = (2.0*p1.y + 1.0*p2.y) / 3.0
	t.x = (1.0*p1.x + 2.0*p2.x) / 3.0
	t.y = (1.0*p1.y + 2.0*p2.y) / 3.0
	// rotation matrix
	u.x = (t.x-s.x)*math.Cos(th) - (t.y-s.y)*math.Sin(th) + s.x
	u.y = (t.x-s.x)*math.Sin(th) + (t.y-s.y)*math.Cos(th) + s.y

	koch(d-1, p1, s)
	fmt.Println(s.x, s.y)
	koch(d-1, s, u)
	fmt.Println(u.x, s.y)
	koch(d-1, u, t)
	fmt.Println(t.x, s.y)
	koch(d-1, t, p2)
}

/************ Koch curve *************/
