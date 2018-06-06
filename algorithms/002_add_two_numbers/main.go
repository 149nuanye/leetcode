package main

import (
	"fmt"
)

// https://leetcode.com/problems/add-two-numbers/description/

// ListNode list node
type ListNode struct {
	Val  int
	Next *ListNode
}

//HeadInsert insert value in list head
func (l *ListNode) HeadInsert(value int) {
	node := &ListNode{Val: value, Next: l.Next}
	l.Next = node
}

//NewList new a list
func NewList(value int) *ListNode {
	l := &ListNode{Val: value, Next: nil}
	return l
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l, tailNode *ListNode
	l1Node := l1
	l2Node := l2
	var carry int
	nodeNum := 0
	for l1Node != nil || l2Node != nil {
		var x, y int
		if l1Node != nil {
			x = l1Node.Val
			l1Node = l1Node.Next
		}
		if l2Node != nil {
			y = l2Node.Val
			l2Node = l2Node.Next
		}
		value := (x + y + carry) % 10
		carry = (x + y + carry) / 10
		nodeNum++
		if nodeNum == 1 {
			tailNode = &ListNode{Val: value, Next: nil}
			l = tailNode
			continue
		}
		tempNode := &ListNode{Val: value, Next: nil}
		tailNode.Next = tempNode
		tailNode = tempNode
	}
	if carry != 0 {
		tempNode := &ListNode{Val: carry, Next: nil}
		tailNode.Next = tempNode
		tailNode = tempNode
	}
	return l
}

func main() {
	l1 := NewList(5)
	l2 := NewList(5)
	l3 := addTwoNumbers(l1, l2)
	for node := l3; node != nil; node = node.Next {
		fmt.Printf("value:%v\n", node.Val)
	}
}
