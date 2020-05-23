package main
import "fmt"
// TreeNode tree node
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0||len(inorder)==0{
		return nil
	}
	valueMap := make(map[int]int)
	for index,value := range inorder{
		valueMap[value] = index
	}
	value := postorder[len(postorder)-1]
	root := &TreeNode{Val: value}
	index := valueMap[value]
	root.Left = buildTree(inorder[:index],postorder[:index])
	if index != len(inorder)-1{
		root.Right = buildTree(inorder[index+1:],postorder[index:len(postorder)-1])
	}
	return root
}

func preOrder(root *TreeNode){
	if root == nil {
		return 
	}
	fmt.Printf("value:%d\n",root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}

func main(){
	inorder := []int{9,3,15,20,7}
	postorder := []int{9,15,7,20,3}
	root := buildTree(inorder,postorder)
	preOrder(root)
}