package main

import (
	"fmt"
)

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodeLength := 0
	for node := head; node != nil; node = node.Next {
		nodeLength++
	}
	if nodeLength < n {
		return head
	}
	delIndex := nodeLength - n + 1
	if delIndex == 1 {
		temp := head
		head = head.Next
		temp.Next = nil
		return head
	}
	index := 1
	for node := head; node.Next != nil; node = node.Next {
		if index+1 == delIndex {
			temp := node.Next
			node.Next = node.Next.Next
			temp.Next = nil
			break
		}
		index++
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
		fmt.Printf("value:%v\n", node.Val)
	}
	l2 := removeNthFromEnd(l1, 2)
	for node := l2; node != nil; node = node.Next {
		fmt.Printf("value:%v\n", node.Val)
	}
}
