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

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			return true
		}
		return false
	}
	sum = sum - root.Val
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

func main() {
	t1 := NewTree(1)
	t1.LeftNode(2)
	t1.RightNode(3)
	t1.RightNode(4)
	fmt.Printf("value:%+v\n", hasPathSum(t1, 8))
}
