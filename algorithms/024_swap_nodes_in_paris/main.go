package main

import (
	"fmt"
)

// https://leetcode.com/problems/swap-nodes-in-pairs/description/

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

func swapPairs(head *ListNode) *ListNode {
	lnode := head
	var preNode *ListNode
	for lnode != nil {
		if lnode.Next == nil {
			break
		}
		tmp:= lnode.Next
		lnode.Next = tmp.Next
		tmp.Next = lnode
		if lnode == head{
			head = tmp
		}
		if preNode != nil {
			preNode.Next = tmp
		}
		preNode = lnode
		lnode = lnode.Next
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
		fmt.Printf("value:%d\t", node.Val)
	}
	fmt.Printf("\n")
	l2 := swapPairs(l1)
	for node := l2; node != nil; node = node.Next {
		fmt.Printf("value:%d\t", node.Val)
	}
	fmt.Printf("\n")
}
