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

func leafValue(preValue int, sum *int, node *TreeNode) {
	if node == nil {
		return
	}
	preValue = preValue*10 + node.Val
	if node.Left == nil && node.Right == nil {
		*sum = *sum + preValue
		return
	}
	if node.Left != nil {
		leafValue(preValue, sum, node.Left)
	}
	if node.Right != nil {
		leafValue(preValue, sum, node.Right)
	}
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	preValue := 0
	leafValue(preValue, &sum, root)
	return sum
}

func main() {
	t1 := NewTree(1)
	t1.LeftNode(2)
	t1.RightNode(3)
	t1.RightNode(4)
	sum := sumNumbers(t1)
	fmt.Printf("sum:%d\n", sum)
}
