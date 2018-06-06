package main

import (
	"fmt"
	"time"
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
	var pre, pre2 *ListNode
	tail := head
	nodeNum := 1
	for ; tail != nil; tail = tail.Next {
		if nodeNum%2 == 0 {
			if pre2 == nil {
				pre.Next = tail.Next
				tail.Next = pre
				head = tail
			} else {
				pre.Next = tail.Next
				pre2.Next = tail
				tail.Next = pre
			}
			temp := tail
			tail = pre
			pre = temp
		}
		pre2 = pre
		pre = tail
		nodeNum++
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
	l2 := swapPairs(l1)
	for node := l2; node != nil; node = node.Next {
		fmt.Printf("value:%v\n", node.Val)
	}
}
