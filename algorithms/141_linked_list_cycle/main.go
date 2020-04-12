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

// https://leetcode-cn.com/problems/linked-list-cycle/solution/huan-xing-lian-biao-by-leetcode/
func hasCycleHash(head *ListNode) bool {
	filter := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := filter[head]; ok {
			return true
		}
		filter[head] = true
		head = head.Next
	}
	return false
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
	}
	return false
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
	fmt.Printf("value:%+v\n", hasCycle(list))
}
