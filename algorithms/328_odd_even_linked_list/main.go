package main

import "fmt"

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

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	node := head
	var preOddNode *ListNode
	var evenHead, preEvenNode *ListNode
	num := 1
	for node != nil {
		if num%2 != 0 {
			if preOddNode != nil {
				preOddNode.Next = node
			}
			preOddNode = node
			node = node.Next
			preOddNode.Next = nil
		} else {
			if evenHead == nil {
				evenHead = node
			} else {
				preEvenNode.Next = node
			}
			preEvenNode = node
			node = node.Next
			preEvenNode.Next = nil
		}
		num++
	}
	preOddNode.Next = evenHead
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
	fmt.Printf("after oddEvenList\n")
	root := oddEvenList(l1)
	for node := root; node != nil; node = node.Next {
		fmt.Printf("value:%v\n", node.Val)
	}
}
