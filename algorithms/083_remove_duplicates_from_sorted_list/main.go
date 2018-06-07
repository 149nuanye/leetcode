package main

import (
	"fmt"
)

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/

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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fakeNode := &ListNode{}
	fakeNode.Next = head
	preNode, tail := fakeNode, head
	for tail != nil {
		for tail.Next != nil && tail.Next.Val == tail.Val {
			tail = tail.Next
		}
		preNode.Next = tail
		preNode = tail
		tail = tail.Next
	}
	return fakeNode.Next
}

func main() {
	l1 := NewList(5)
	l1.HeadInsert(3)
	l1.HeadInsert(3)
	l1.HeadInsert(1)
	l1.HeadInsert(1)
	for node := l1; node != nil; node = node.Next {
		fmt.Printf("value:%v\n", node.Val)
	}
	l2 := deleteDuplicates(l1)
	for node := l2; node != nil; node = node.Next {
		fmt.Printf("end value:%v\n", node.Val)
	}
}
