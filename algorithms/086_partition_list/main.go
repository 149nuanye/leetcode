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

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	fakeHead := &ListNode{}
	tailNode := fakeHead

	rightHead := &ListNode{}
	rightCur := rightHead

	curNode := head
	for curNode != nil {
		if curNode.Val < x {
			tailNode.Next = curNode
			tailNode = curNode
		} else {
			rightCur.Next = curNode
			rightCur = curNode
		}
		curNode = curNode.Next
	}
	rightCur.Next = nil
	tailNode.Next = rightHead.Next
	return fakeHead.Next
}

func main() {
	l1 := NewList(2)
	l1.HeadInsert(5)
	l1.HeadInsert(2)
	l1.HeadInsert(3)
	l1.HeadInsert(4)
	l1.HeadInsert(1)
	for node := l1; node != nil; node = node.Next {
		fmt.Printf("value:%v\n", node.Val)
	}
	l2 := partition(l1, 3)
	for node := l2; node != nil; node = node.Next {
		fmt.Printf("end value:%v\n", node.Val)
	}
}
