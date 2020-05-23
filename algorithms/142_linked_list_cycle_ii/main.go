package main

import "fmt"

// ListNode list node
type ListNode struct {
	Val  int
	Next *ListNode
}

// Insert insert
func (l *ListNode) Insert(node *ListNode) {
	if l.Next == nil {
		l.Next = node
	} else {
		node.Next = l.Next
		l.Next = node
	}
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = head
	for fast != nil {
		if fast == slow {
			return fast
		}
		fast = fast.Next
		slow = slow.Next
	}
	return nil
}

func main() {
	list := &ListNode{Val: 3}
	node4 := &ListNode{Val: -4}
	list.Insert(node4)
	node0 := &ListNode{Val: 0}
	list.Insert(node0)
	node2 := &ListNode{Val: 2}
	node4.Next = node2
	list.Insert(node2)
	node := detectCycle(list)
	fmt.Printf("value:%d\n", node.Val)
}
