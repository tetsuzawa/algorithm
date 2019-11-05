package main

import "fmt"

const (
	MAX = 100005
	NIL = -1
)

type Node struct {
	Parent int
	Left   int
	Right  int
}

var T = make([]Node, MAX)
var n int
var D []int

func Print(u int) {
	fmt.Printf("node %v: ", u)
	fmt.Printf("parent = %v, ", T[u].Parent)
	fmt.Printf("depth = %v, ", D[u])

	if T[u].Parent == NIL {
		fmt.Print("root, ")
	} else if T[u].Left == NIL {
		fmt.Print("leaf, ")
	} else {
		fmt.Print("internal node , ")
	}

	fmt.Print("[")
	c := T[u].Left
	for i := 0; c != NIL; i++ {
		if i != 0 {
			fmt.Print(", ")
		}
		fmt.Println("]")
		c = T[c].Right
	}

}

func rec(u, p int) {
	D[u] = p
	if T[u].Right != -1 {
		rec(T[u].Right, p)
	}
	if T[u].Left != -1 {
		rec(T[u].Left, p+1)
	}
}

func main() {

}

func GetDepth(T []Node, u int) int {
	var d int
	for T[u].Parent != -1 {
		u = T[u].Parent
		d++
	}
	return d
}

func SetDepth(T []Node, D []int, u, p int) {
	D[u] = p
	if T[u].Right != -1 {
		SetDepth(T, D, T[u].Right, p)
	}
	if T[u].Left != -1 {
		SetDepth(T, D, T[u].Left, p+1)
	}
}

func PrintChildren(T []Node, u int) {
	c := T[u].Left
	for c != -1 {
		fmt.Println(c)
		c = T[c].Right
	}
}
