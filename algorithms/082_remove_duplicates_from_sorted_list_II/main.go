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

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	var preNode *ListNode
// 	tail := head
// 	dupTag := 0
// 	for tail != nil && tail.Next != nil {
// 		if tail.Val == tail.Next.Val {
// 			dupTag = 1
// 			nextNode := tail.Next
// 			tail.Next = nextNode.Next
// 			nextNode.Next = nil
// 			continue
// 		}
// 		if dupTag == 1 {
// 			dupTag = 0
// 			if preNode == nil {
// 				head = tail.Next
// 				tail.Next = nil
// 				tail = head
// 				continue
// 			}
// 			preNode.Next = tail.Next
// 			tail.Next = nil
// 			tail = preNode.Next
// 			continue
// 		}
// 		preNode = tail
// 		tail = tail.Next
// 	}
// 	if dupTag == 1 {
// 		if preNode == nil {
// 			head = tail.Next
// 			tail.Next = nil
// 			return head
// 		}
// 		preNode.Next = nil
// 	}
// 	return head
// }

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fakeHead := &ListNode{}
	fakeHead.Next = head
	preNode, tail := fakeHead, head
	for tail != nil {
		for tail.Next != nil && tail.Val == tail.Next.Val {
			tail = tail.Next
		}
		if tail == preNode.Next {
			preNode = tail
		} else {
			preNode.Next = tail.Next
		}
		tail = tail.Next
	}
	return fakeHead.Next
}

func main() {
	l1 := NewList(5)
	l1.HeadInsert(1)
	l1.HeadInsert(1)

	for node := l1; node != nil; node = node.Next {
		fmt.Printf("value:%v\n", node.Val)
	}
	l2 := deleteDuplicates(l1)
	for node := l2; node != nil; node = node.Next {
		fmt.Printf("end value:%v\n", node.Val)
	}
}
