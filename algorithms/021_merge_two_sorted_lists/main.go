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

// first get list length,then get location of the node to be deleted
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3, l3Tail *ListNode
	fmt.Println(l3)
	head1 := l1
	head2 := l2
	for head1 != nil || head2 != nil {
		var node *ListNode
		if head1 == nil {
			node = &ListNode{Val: head2.Val, Next: nil}
			head2 = head2.Next
		} else if head2 == nil {
			node = &ListNode{Val: head1.Val, Next: nil}
			head1 = head1.Next
		} else if head1.Val < head2.Val {
			node = &ListNode{Val: head1.Val, Next: nil}
			head1 = head1.Next
		} else {
			node = &ListNode{Val: head2.Val, Next: nil}
			head2 = head2.Next
		}
		if l3 == nil {
			l3 = node
			l3Tail = l3
			continue
		}
		l3Tail.Next = node
		l3Tail = node
	}
	return l3
}

func main() {
	l1 := NewList(5)
	l1.HeadInsert(4)
	l1.HeadInsert(3)
	l1.HeadInsert(2)
	l1.HeadInsert(1)
	for node := l1; node != nil; node = node.Next {
		fmt.Printf("value:%v\n", node.Val)
	}
	l2 := NewList(5)
	l2.HeadInsert(4)
	l3 := mergeTwoLists(l1, l2)
	for node := l3; node != nil; node = node.Next {
		fmt.Printf("value:%v\n", node.Val)
	}
}
