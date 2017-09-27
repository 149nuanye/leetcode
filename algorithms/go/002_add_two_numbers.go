package main

import (
	"fmt"
)

//desc https://leetcode.com/problems/add-two-numbers/description/

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//ListInit init list
func ListInit(value int) *ListNode {
	list := ListNode{Val: value}
	list.Next = nil
	return &list
}

//AddNode add node
func AddNode(list *ListNode, value int) *ListNode {
	tmpNode := &ListNode{Val: value}
	tmpNode.Next = list
	list = tmpNode
	return list
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp1, tmp2 := l1, l2
	tmp3 := tmp1
	tmp := 0
	for ; tmp1 != nil && tmp2 != nil; tmp1, tmp2 = tmp1.Next, tmp2.Next {
		if tmp1.Val+tmp2.Val+tmp >= 10 {
			tmp1.Val = tmp1.Val + tmp2.Val + tmp - 10
			tmp = 1
		} else {
			tmp1.Val = tmp1.Val + tmp2.Val + tmp
			tmp = 0
		}
		tmp3 = tmp1
		// fmt.Printf("hello,tmp3:%v,tmp1:%v\n", tmp3, tmp1)
	}
	if tmp1 != nil {
		for ; tmp1 != nil; tmp1 = tmp1.Next {
			if tmp1.Val+tmp >= 10 {
				tmp1.Val = tmp1.Val + tmp - 10
				tmp = 1
			} else {
				tmp1.Val = tmp1.Val + tmp
				tmp = 0
			}
			tmp3 = tmp1
		}
	}
	if tmp2 != nil {
		tmp3.Next = tmp2
		for ; tmp2 != nil; tmp2 = tmp2.Next {
			if tmp2.Val+tmp >= 10 {
				tmp2.Val = tmp2.Val + tmp - 10
				tmp = 1
			} else {
				tmp2.Val = tmp2.Val + tmp
				tmp = 0
			}
			tmp3 = tmp2
		}
	}
	if tmp == 1 {
		tmpNode := &ListNode{Val: tmp}
		tmp3.Next = tmpNode
	}
	return l1
}
func main() {
	list1 := ListInit(9)
	// arr1 := []int{4, 2}
	// for index := range arr1 {
	// 	list1 = AddNode(list1, arr1[index])
	// }
	for tmp := list1; tmp != nil; tmp = tmp.Next {
		fmt.Printf("%d\t", tmp.Val)
	}
	fmt.Printf("\n")
	list2 := ListInit(9)
	list2 = AddNode(list2, 9)
	// arr2 := []int{6, 8, 4}
	// for index := range arr2 {
	// 	list2 = AddNode(list2, arr2[index])
	// }
	for tmp := list2; tmp != nil; tmp = tmp.Next {
		fmt.Printf("%d\t", tmp.Val)
	}
	fmt.Printf("\n")
	list3 := addTwoNumbers(list2, list1)
	for tmp := list3; tmp != nil; tmp = tmp.Next {
		fmt.Printf("%d\t", tmp.Val)
	}
	fmt.Printf("\n")
}
