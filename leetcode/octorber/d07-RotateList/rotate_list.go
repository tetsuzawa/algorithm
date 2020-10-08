package d07_RotateList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	node := head
	var a []int
	for node != nil {
		a = append(a, node.Val)
		node = node.Next
	}
	if a == nil {
		return nil
	}
	fmt.Println(a)
	for i := 0; i < k%len(a); i++ {
		a = append([]int{a[len(a)-1]}, a[:len(a)-1]...)

	}
	fmt.Println(a)
	node = &ListNode{Val: a[0]}
	rootNode := node
	a = a[1:]
	for len(a) != 0 {
		node.Next = &ListNode{Val: a[0]}
		a = a[1:]
		node = node.Next
	}
	return rootNode
}
