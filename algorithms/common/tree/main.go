package main

import "fmt"

// Queue queue
type Queue struct {
	Data []*TreeNode
}

// NewQueue new queue
func NewQueue() *Queue {
	queue := Queue{Data: make([]*TreeNode, 0, 10)}
	return &queue
}

// Push push data
func (queue *Queue) Push(node *TreeNode) {
	queue.Data = append(queue.Data, node)
}

// Pop pop data
func (queue *Queue) Pop() *TreeNode {
	if len(queue.Data) == 0 {
		return nil
	}
	node := queue.Data[0]
	queue.Data = queue.Data[1:]
	return node
}

// Length data length
func (queue *Queue) Length() int {
	return len(queue.Data)
}

// Stack stack
type Stack struct {
	Data []*TreeNode
}

// NewStack new Stack
func NewStack() *Stack {
	stack := Stack{Data: make([]*TreeNode, 0, 10)}
	return &stack
}

// Push push data
func (stack *Stack) Push(node *TreeNode) {
	stack.Data = append(stack.Data, node)
}

// Pop pop data
func (stack *Stack) Pop() *TreeNode {
	if len(stack.Data) == 0 {
		return nil
	}
	length := len(stack.Data) - 1
	node := stack.Data[length]
	stack.Data = stack.Data[:length]
	return node
}

// Length stack length
func (stack *Stack) Length() int {
	return len(stack.Data)
}

func traversalOrder(tree *Tree) []int {
	if tree.root == nil {
		return nil
	}
	data := make([]int, 0, 10)
	queue := NewQueue()
	queue.Push(tree.root)
	for queue.Length() > 0 {
		node := queue.Pop()
		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
		}
		data = append(data, node.Data)
	}
	return data
}

func preOrderRecursive(tree *TreeNode) []int {
	if tree == nil {
		return nil
	}
	data := make([]int, 0, 10)
	data = append(data, tree.Data)
	if tree.Left != nil {
		tmp := preOrderRecursive(tree.Left)
		if tmp != nil {
			data = append(data, tmp...)
		}
	}
	if tree.Right != nil {
		tmp := preOrderRecursive(tree.Right)
		if tmp != nil {
			data = append(data, tmp...)
		}
	}
	return data
}

func preOrder(tree *Tree) []int {
	if tree == nil {
		return nil
	}
	node := tree.root
	stack := NewStack()
	data := make([]int, 0, 10)
	for node != nil || stack.Length() > 0 {
		for node != nil {
			stack.Push(node)
			data = append(data, node.Data)
			node = node.Left
		}
		tmpNode := stack.Pop()
		node = tmpNode.Right
	}
	return data
}

func inOrder(tree *Tree) []int {
	if tree == nil {
		return nil
	}
	node := tree.root
	stack := NewStack()
	data := make([]int, 0, 10)
	for node != nil || stack.Length() > 0 {
		for node != nil {
			stack.Push(node)
			node = node.Left
		}
		tmpNode := stack.Pop()
		data = append(data, tmpNode.Data)
		node = tmpNode.Right
	}
	return data
}

func postOrder(tree *Tree) []int {
	if tree == nil {
		return nil
	}
	node := tree.root
	stack := NewStack()
	data := make([]int, 0, 10)
	nodeMap := make(map[*TreeNode]struct{})
	for node != nil || stack.Length() > 0 {
		for node != nil {
			stack.Push(node)
			node = node.Left
		}
		tmpNode := stack.Pop()
		if tmpNode.Right == nil {
			data = append(data, tmpNode.Data)
			continue
		}
		if _, ok := nodeMap[tmpNode]; ok {
			data = append(data, tmpNode.Data)
			continue
		}
		nodeMap[tmpNode] = struct{}{}
		stack.Push(tmpNode)
		node = tmpNode.Right
	}
	return data
}

func main() {
	tree := &Tree{}
	tree.Add(50)
	tree.Add(45)
	tree.Add(40)
	tree.Add(48)
	tree.Add(51)
	tree.Add(61)
	tree.Add(71)
	// data := traversalOrder(tree)
	data := preOrderRecursive(tree.root)
	fmt.Printf("preOrderRecursive data:%+v\n", data)
	data = postOrder(tree)
	fmt.Printf("postOrder data:%+v\n", data)
}
