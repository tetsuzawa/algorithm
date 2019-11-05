package tree

import "fmt"

// Left child, right sibling, representation.
type Node struct {
	Parent int
	Left   int
	Right  int
}

// Get depth of nodes.
// O(nh).
func GetDepth(T []Node, u int) int {
	var d int
	for T[u].Parent != -1 {
		u = T[u].Parent
		d++
	}
	return d
}

// Get depth of nodes (recursive).
// O(n). Faster algorithm.
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
