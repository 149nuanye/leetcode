package main

import "fmt"

// TreeNode treeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
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

// https://leetcode-cn.com/problems/symmetric-tree/solution/dui-cheng-er-cha-shu-by-leetcode/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val == t2.Val {
		return isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
	}
	return false
}

func main() {
	t1 := NewTree(1)
	t1.LeftNode(2)
	t1.RightNode(2)
	fmt.Printf("isSymmetric:%+v\n", isSymmetric(t1))
}
