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
	var lHead, rhead *ListNode
	var lCur, rCur *ListNode
	curNode := head
	for curNode != nil {
		tmp := curNode.Next
		if curNode.Val >= x {
			curNode.Next = nil
			if rhead == nil {
				rhead = curNode
				rCur = curNode
			} else {
				rCur.Next = curNode
				rCur = curNode
			}
			curNode = tmp
			continue
		}

		if lHead == nil {
			lHead = curNode
			lCur = curNode
		} else {
			lCur.Next = curNode
			lCur = curNode
		}
		curNode = tmp
	}
	if lHead == nil {
		return rhead
	}
	if rhead != nil {
		lCur.Next = rhead
	}
	return lHead
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
	l2 := partitionDemo(l1, 3)
	for node := l2; node != nil; node = node.Next {
		fmt.Printf("end value:%v\n", node.Val)
	}
}
