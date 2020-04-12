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

//NewList new a list
func NewList(value int) *ListNode {
	l := &ListNode{Val: value, Next: nil}
	return l
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry :=0
	var newList *ListNode
	var resultList *ListNode
	for(l1!=nil&&l2!=nil){
		tmpVal := (l1.Val + l2.Val+carry)%10 
		carry = (l1.Val + l2.Val+carry)/10 
		newNode := ListNode{Val: tmpVal, Next: nil}
		if newList == nil{
			newList = &newNode
			resultList = newList 
		} else {
			newList.Next = &newNode
			newList=&newNode
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	for (l1 != nil){
		tmpVal := (l1.Val + carry)%10 
		carry = (l1.Val + carry)/10 
		newNode:=ListNode{Val: tmpVal, Next: nil}
		newList.Next = &newNode
		newList = &newNode
		l1 = l1.Next
	}
	for (l2!= nil){
		tmpVal := (l2.Val + carry)%10 
		carry = (l2.Val + carry)/10 
		newNode:=ListNode{Val: tmpVal, Next: nil}
		newList.Next = &newNode
		newList = &newNode
		l2 = l2.Next
	}
	if carry > 0 {
		newNode:=ListNode{Val: carry, Next: nil}
		newList.Next = &newNode
		newList = &newNode
	}
	return resultList 
}

func main() {
	l1 := NewList(5)
	l2 := NewList(5)
	l3 := addTwoNumbers(l1, l2)
	for node := l3; node != nil; node = node.Next {
		fmt.Printf("value:%v\n", node.Val)
	}
}
