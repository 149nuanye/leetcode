package main

import (
	"fmt"
)

// https://leetcode.com/problems/merge-two-sorted-lists/description/

// ListNode list node
type ListNode struct {
	Val  int
	Next *ListNode
}

//HeadInsert insert value in list head
func (l *ListNode) HeadInsert(value int) {
	node := &ListNode{Val: l.Val, Next: l.Next}
	l.Next = node
	l.Val = value
}

// NewList new a list
func NewList(value int) *ListNode {
	l := &ListNode{Val: value, Next: nil}
	return l
}

//递归
func mergeTwoListsNew(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var node *ListNode
	if l1.Val > l2.Val {
		node = l2
		node.Next = mergeTwoListsNew(l1, l2.Next)
	} else {
		node = l1
		node.Next = mergeTwoListsNew(l1.Next, l2)
	}
	return node
}

//循环
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head, tail *ListNode
	node1 := l1
	node2 := l2
	for node1 != nil && node2 != nil {
		if node1.Val > node2.Val {
			if head == nil {
				head = node2
				tail = node2
			} else {
				tail.Next = node2
				tail = node2
			}
			node2 = node2.Next
		} else {
			if head == nil {
				head = node1
				tail = node1
			} else {
				tail.Next = node1
				tail = node1
			}
			node1 = node1.Next
		}
	}
	if node1 != nil {
		tail.Next = node1
	}
	if node2 != nil {
		tail.Next = node2
	}
	return head
}

func main() {
	l1 := NewList(5)
	l1.HeadInsert(4)
	l1.HeadInsert(3)
	l1.HeadInsert(2)
	l1.HeadInsert(1)
	for node := l1; node != nil; node = node.Next {
		fmt.Printf("node1 value:%v\n", node.Val)
	}
	l2 := NewList(5)
	l2.HeadInsert(4)
	l3 := mergeTwoListsNew(l1, l2)
	for node := l3; node != nil; node = node.Next {
		fmt.Printf("merger value:%v\n", node.Val)
	}
}
