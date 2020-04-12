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
// 官方解答：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/solution/shan-chu-lian-biao-de-dao-shu-di-nge-jie-dian-by-l/
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

func removeNthFromEndOne(head *ListNode, n int) *ListNode {
	lNode := head
	rNode := head
	for index := 0; index < n; index++{
		if rNode == nil {
			return head
		}
		rNode = rNode.Next
	}
	// 删除第一个节点
	if rNode == nil {
		tmp := head
		head = head.Next
		tmp.Next = nil 
		return head
	}
	prNode := lNode
	for rNode != nil {
		prNode = lNode
		lNode = lNode.Next
		rNode = rNode.Next
	}
	prNode.Next = lNode.Next
	lNode.Next = nil 
	return head
}

func main() {
	l1 := NewList(5)
	l1.HeadInsert(4)
	l1.HeadInsert(3)
	l1.HeadInsert(2)
	l1.HeadInsert(1)
	fmt.Printf("value:")
	for node := l1; node != nil; node = node.Next {
		fmt.Printf("%v\t", node.Val)
	}
	fmt.Printf("\n")
	l2 := removeNthFromEndOne(l1, 2)
	fmt.Printf("value:")
	for node := l2; node != nil; node = node.Next {
		fmt.Printf("%v\t", node.Val)
	}
	fmt.Printf("\n")
}
