package dataStructure

// Left child, right sibling, representation.
type Node struct {
	Parent int
	Left   int
	Right  int
}

// Get depth of nodes.
func (n Node) GetDepth(T []Node, u int) int {
	var d int
	for T[u].Parent != -1 {
		u = T[u].Parent
		d++
	}
	return d
}

// Get depth of nodes (recursive).
func (n Node) GetDepthR(T []Node, D []int, u int) int {
	D
}
