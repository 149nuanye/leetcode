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

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	for i := head; i != nil; i = i.Next {
		for j := head; j != i; j = j.Next {
			if i.Val < j.Val {
				i.Val, j.Val = j.Val, i.Val
			}
		}
	}
	return head
}

func main() {
	l1 := NewList(2)
	l1.HeadInsert(5)
	l1.HeadInsert(2)
	for node := l1; node != nil; node = node.Next {
		fmt.Printf("value:%v\n", node.Val)
	}
	l2 := insertionSortList(l1)
	for node := l2; node != nil; node = node.Next {
		fmt.Printf("end value:%v\n", node.Val)
	}
}
