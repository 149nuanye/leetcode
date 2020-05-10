package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// InsertNode insert value in list head
func (l *ListNode) InsertNode(node *ListNode) {
	node.Next = l.Next
	l.Next = node
}

// NewList new a list
func NewList(value int) *ListNode {
	l := &ListNode{Val: value, Next: nil}
	return l
}

// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/solution/xiang-jiao-lian-biao-by-leetcode/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeA := headA
	nodeB := headB
	tagA, tagB := false, false
	for nodeA != nil && nodeB != nil {
		if nodeA == nodeB {
			return nodeA
		}
		nodeA = nodeA.Next
		nodeB = nodeB.Next
		if nodeA == nil && !tagA {
			nodeA = headB
			tagA = true
		}
		if nodeB == nil && !tagB {
			nodeB = headA
			tagB = true
		}
	}
	return nil
}

func main() {
	l1 := NewList(5)
	l2 := NewList(6)
	node := &ListNode{Val: 4, Next: nil}
	l1.InsertNode(node)
	node = &ListNode{Val: 3, Next: nil}
	l1.InsertNode(node)
	l2.InsertNode(node)
	node = &ListNode{Val: 2, Next: nil}
	l1.InsertNode(node)
	l2.InsertNode(node)
	node = &ListNode{Val: 1, Next: nil}
	l1.InsertNode(node)
	l2.InsertNode(node)

	node = getIntersectionNode(l1, l2)
	for node != nil {
		fmt.Printf("value:%+v\n", node.Val)
		node = node.Next
	}
}
