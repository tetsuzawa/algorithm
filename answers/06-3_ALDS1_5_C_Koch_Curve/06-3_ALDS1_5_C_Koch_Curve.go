package _6_3_ALDS1_5_C_Koch_Curve

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

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
	fmt.Println(u.x, u.y)
	koch(d-1, u, t)
	fmt.Println(t.x, t.y)
	koch(d-1, t, p2)
}

func MainKoch() {
	sc := bufio.NewScanner(os.Stdin)
	var p1 = Point{
		x: 0,
		y: 0,
	}
	var p2= Point{
		x: 100,
		y: 0,
	}
	sc.Scan()
	d, _ := strconv.Atoi(sc.Text())

	fmt.Println(p1.x, p1.y)
	koch(d, p1, p2)
	fmt.Println(p2.x, p2.y)
}
