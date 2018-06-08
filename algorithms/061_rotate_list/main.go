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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curNode := head
	tailNode := head
	nodeLength := 0
	for curNode != nil {
		tailNode = curNode
		curNode = curNode.Next
		nodeLength++
	}
	index := nodeLength - k%nodeLength
	if index == nodeLength {
		return head
	}
	cutNode := head
	for i := 1; i < index; i++ {
		cutNode = cutNode.Next
	}
	tempNode := cutNode.Next
	cutNode.Next = nil
	tailNode.Next = head
	head = tempNode
	return head
}

func main() {
	l1 := NewList(1)
	l1.HeadInsert(2)
	for node := l1; node != nil; node = node.Next {
		fmt.Printf("value:%v\n", node.Val)
	}
	l2 := rotateRight(l1, 2)
	for node := l2; node != nil; node = node.Next {
		fmt.Printf("value:%v\n", node.Val)
	}
}
