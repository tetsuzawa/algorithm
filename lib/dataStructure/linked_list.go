package dataStructure

type Node struct {
	Key  int
	Prev *Node
	Next *Node
}

var null *Node

func init() {
	null := new(Node)
	null.Next = null
	null.Next = null
}

func insert(key int)  {
	var x = new(*Node)
	x.Key = key
	x.Next = null.Next
	null.Next.Prev = x
	null.Next = x
	x.Prev = null

}



