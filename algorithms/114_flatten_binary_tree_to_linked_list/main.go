package main

import "fmt"

// TreeNode tree node
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//NewTree new tree
func NewTree(val int) *TreeNode {
	tree := TreeNode{Val: val}
	return &tree
}

// LeftNode left node
func (tree *TreeNode) LeftNode(val int) {
	node := tree
	parentNode := node
	for node != nil {
		parentNode = node
		node = node.Left
	}
	parentNode.Left = &TreeNode{Val: val}
}

// RightNode right node
func (tree *TreeNode) RightNode(val int) {
	node := tree
	parentNode := node
	for node != nil {
		parentNode = node
		node = node.Right
	}
	parentNode.Right = &TreeNode{Val: val}
}
// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/solution/114-er-cha-shu-zhan-kai-wei-lian-biao-by-ming-zhi-/
func flatten(root *TreeNode)  {
	if root == nil {
		return 
	}
	flatten(root.Left)
	flatten(root.Right)
	tmp := root.Right
	root.Right = root.Left
	root.Left = nil 
	for root.Right != nil {
		root = root.Right
	}
	root.Right = tmp
}

func main(){
	t1 := NewTree(1)
	t1.LeftNode(2)
	t1.RightNode(3)
	t1.RightNode(4)
	for t1 != nil {
		fmt.Printf("value:%d\n",t1.Val)
		t1 = t1.Right
	}
}