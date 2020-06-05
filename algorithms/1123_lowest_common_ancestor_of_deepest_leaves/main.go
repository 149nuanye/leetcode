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

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-deepest-leaves/solution/dfsjian-dan-jie-ti-si-lu-yi-kan-jiu-dong-by-xiaora/
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	ld := depth(root.Left)
	rd := depth(root.Right)
	// fmt.Printf("value:%d,ld:%d,rd:%d\n", root.Val, ld, rd)
	if ld == rd {
		return root
	}
	if ld > rd {
		return lcaDeepestLeaves(root.Left)
	}
	return lcaDeepestLeaves(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depthV := 1
	if root.Left != nil {
		depthV = depth(root.Left) + 1
	}
	if root.Right != nil {
		tmp := depth(root.Right) + 1
		if tmp > depthV {
			depthV = tmp
		}
	}
	return depthV
}

func main() {
	t1 := NewTree(1)
	t1.LeftNode(2)
	t1.RightNode(3)
	t1.RightNode(4)
	node := lcaDeepestLeaves(t1)
	fmt.Printf("value:%d\n", node.Val)
}
