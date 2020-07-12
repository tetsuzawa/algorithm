package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// find height
// O(n)

const (
	MAX = 10000
	NIL = -1
)

type Node struct {
	parent int
	left   int
	right  int
}

var n int
var T = make([]Node, MAX)
var D = make([]int, MAX)
var H = make([]int, MAX)

func setDepth(u, d int) {
	if u == NIL {
		return
	}
	D[u] = d
	setDepth(T[u].left, d+1)
	setDepth(T[u].right, d+1)
}

func setHeight(u int) int {
	var h1, h2 int
	if T[u].left != NIL {
		h1 = setHeight(T[u].left) + 1
	}
	if T[u].right != NIL {
		h2 = setHeight(T[u].right) + 1
	}
	if h1 > h2 {
		H[u] = h1
	} else {
		H[u] = h2
	}
	return H[u]
}

// return sibling of node u
func getSibling(u int) int {
	if T[u].parent == NIL {
		return NIL
	}
	if T[T[u].parent].left != u && T[T[u].parent].left != NIL {
		return T[T[u].parent].left
	}
	if T[T[u].parent].right != u && T[T[u].parent].right != NIL {
		return T[T[u].parent].right
	}
	return NIL
}

func Print(u int) {
	fmt.Printf("node %d: ", u)
	fmt.Printf("parent = %d, ", T[u].parent)
	fmt.Printf("sibling = %d, ", getSibling(u))
	var deg int
	if T[u].left != NIL {
		deg++
	}
	if T[u].right != NIL {
		deg++
	}
	fmt.Printf("degree = %d, ", deg)
	fmt.Printf("depth = %d, ", D[u])
	fmt.Printf("height = %d, ", H[u])

	if T[u].parent == NIL {
		fmt.Println("root")
	} else if T[u].left == NIL && T[u].right == NIL {
		fmt.Println("leaf")
	} else {
		fmt.Println("internal node")
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var v, l, r, root int
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		T[i].parent = NIL
	}
	for i := 0; i < n; i++ {
		sc.Scan()
		ss := strings.Split(sc.Text(), " ")
		v, _ = strconv.Atoi(ss[0])
		l, _ = strconv.Atoi(ss[1])
		r, _ = strconv.Atoi(ss[2])
		T[v].left = l
		T[v].right = r
		if l != NIL {
			T[l].parent = v
		}
		if r != NIL {
			T[r].parent = v
		}
	}
	for i := 0; i < n; i++ {
		if T[i].parent == NIL {
			root = i
		}
	}
	setDepth(root, 0)
	setHeight(root)
	for i := 0; i < n; i++ {
		Print(i)
	}
}
