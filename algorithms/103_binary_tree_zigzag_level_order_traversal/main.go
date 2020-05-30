package main

import (
	"fmt"
)

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

// Queue queue
type Queue struct {
	Data []*TreeNode
}

// NewQueue new queue
func NewQueue() *Queue {
	queue := &Queue{Data: make([]*TreeNode, 0, 100)}
	return queue
}

func (q *Queue) Push(b *TreeNode) {
	q.Data = append(q.Data, b)
}

func (q *Queue) Pop() *TreeNode {
	b := q.Data[0]
	q.Data = q.Data[1:]
	return b
}

func (q *Queue) Length() int {
	return len(q.Data)
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	order := make([][]int, 0, 10)
	if root == nil {
		return order
	}
	q := NewQueue()
	q.Push(root)
	level := 0
	for q.Length() > 0 {
		nodeList := make([]*TreeNode, 0, q.Length())
		nodeVals := make([]int, 0, q.Length())
		for q.Length() > 0 {
			node := q.Pop()
			nodeVals = append(nodeVals, node.Val)
			if level%2 == 0 {
				if node.Left != nil {
					nodeList = append(nodeList, node.Left)
				}
				if node.Right != nil {
					nodeList = append(nodeList, node.Right)
				}
			} else {
				if node.Right != nil {
					nodeList = append(nodeList, node.Right)
				}
				if node.Left != nil {
					nodeList = append(nodeList, node.Left)
				}
			}
		}
		fmt.Printf("level:%+v, nodeVals:%+v\n", level, nodeVals)
		order = append(order, nodeVals)
		for index := len(nodeList) - 1; index >= 0; index-- {
			q.Push(nodeList[index])
		}
		level++
	}
	return order
}

func main() {
	t1 := NewTree(1)
	t1.LeftNode(2)
	t1.RightNode(3)
	t1.RightNode(4)
	t1.LeftNode(5)
	order := zigzagLevelOrder(t1)
	fmt.Printf("order:%+v\n", order)
}
