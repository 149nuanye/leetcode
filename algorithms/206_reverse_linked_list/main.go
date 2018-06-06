package main

import (
	"fmt"
)

// https://leetcode.com/problems/reverse-linked-list/description/

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

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nextNode := head.Next
	currNode := head
	currNode.Next = nil
	var tempNode *ListNode
	for nextNode != nil {
		tempNode = nextNode.Next
		nextNode.Next = currNode
		currNode = nextNode
		nextNode = tempNode
	}
	return currNode
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
	l2 := reverseList(l1)
	for node := l2; node != nil; node = node.Next {
		fmt.Printf("value:%v\n", node.Val)
	}
}
