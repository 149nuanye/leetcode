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
	tail := head
	for tail != nil && tail.Next != nil {
		if tail.Val == tail.Next.Val {
			nextNode := tail.Next
			tail.Next = nextNode.Next
			nextNode.Next = nil
			continue
		}
		tail = tail.Next
	}
	return head
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
		fmt.Printf("value:%v\n", node.Val)
	}
}
