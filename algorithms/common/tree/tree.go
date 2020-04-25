package main

// TreeNode tree node
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

// Tree tree
type Tree struct {
	root *TreeNode
}

//Add add node
func (tree *Tree) Add(data int) {
	newNode := &TreeNode{Data: data}
	if tree.root == nil {
		tree.root = newNode
		return
	}
	parentNode := tree.root
	node := tree.root
	for node != nil {
		parentNode = node
		if node.Data > data {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	if parentNode.Data > data {
		parentNode.Left = &TreeNode{Data: data}
	} else {
		parentNode.Right = &TreeNode{Data: data}
	}
}
