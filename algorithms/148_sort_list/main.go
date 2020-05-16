package main

import (
	"fmt"
)

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

// 递归
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	next := slow.Next
	slow.Next = nil
	lnode := sortList(head)
	rnode := sortList(next)
	return mergeTwoLists(lnode, rnode)
}

func main() {
	l1 := NewList(5)
	l1.HeadInsert(4)
	l1.HeadInsert(1)
	l1.HeadInsert(2)
	l1.HeadInsert(3)
	for node := l1; node != nil; node = node.Next {
		fmt.Printf("node1 value:%v\n", node.Val)
	}
	sortedNode := sortList(l1)
	for node := sortedNode; node != nil; node = node.Next {
		fmt.Printf("sort value:%v\n", node.Val)
	}
}
