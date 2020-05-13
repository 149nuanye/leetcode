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
func (tree *TreeNode) LeftNode(node *TreeNode) {
	tmp := tree
	parentNode := tmp
	for tmp != nil {
		parentNode = tmp
		tmp = tmp.Left
	}
	parentNode.Left = node
}

// RightNode right node
func (tree *TreeNode) RightNode(node *TreeNode) {
	tmp := tree
	parentNode := tmp
	for tmp != nil {
		parentNode = tmp
		tmp = tmp.Right
	}
	parentNode.Right = node
}

func dfs(parentNode, node *TreeNode, parentInfo map[*TreeNode]*TreeNode) {
	if node == nil {
		return
	}
	parentInfo[node] = parentNode
	dfs(node, node.Left, parentInfo)
	dfs(node, node.Right, parentInfo)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	parentInfo := make(map[*TreeNode]*TreeNode)
	dfs(nil, root, parentInfo)
	fmt.Printf("parentInfo:%+v\n", parentInfo)

	pParent := make(map[*TreeNode]struct{})
	for p != nil {
		pParent[p] = struct{}{}
		p = parentInfo[p]
	}
	fmt.Printf("pParent:%+v\n", pParent)
	for q != nil {
		if _, ok := pParent[q]; ok {
			return q
		}
		q = parentInfo[q]
	}
	return nil
}

func main() {
	t1 := NewTree(1)
	node2 := TreeNode{Val: 2}
	t1.LeftNode(&node2)
	node3 := TreeNode{Val: 3}
	t1.RightNode(&node3)
	node4 := TreeNode{Val: 4}
	t1.RightNode(&node4)
	node := lowestCommonAncestor(t1, &node2, &node4)
	fmt.Printf("Node:%+v\n", node)
}
