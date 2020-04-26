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

func maxTreeDepth(node *TreeNode, parentMax int, treeMax *int) {
	if node == nil {
		return
	}
	parentMax++
	if parentMax > *treeMax {
		*treeMax = parentMax
	}
	if node.Left != nil {
		maxTreeDepth(node.Left, parentMax, treeMax)
	}
	if node.Right != nil {
		maxTreeDepth(node.Right, parentMax, treeMax)
	}
}

func maxDepth(root *TreeNode) int {
	max := 0
	parentMax := 0
	maxTreeDepth(root, parentMax, &max)
	return max
}

func main() {
	t1 := NewTree(1)
	t1.LeftNode(2)
	t1.RightNode(3)
	t1.RightNode(4)
	fmt.Printf("maxDepth:%+v\n", maxDepth(t1))
}
