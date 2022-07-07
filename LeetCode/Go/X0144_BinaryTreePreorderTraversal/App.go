package X0144_BinaryTreePreorderTraversal

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Main() {
	root := &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	output := preorderTraversal(root)
	result, _ := json.Marshal(output)
	fmt.Println("Result Binary Tree Preorder Traversal: " + string(result))

}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		result = append(result, root.Val)
		result = append(result, preorderTraversal(root.Left)...)
		result = append(result, preorderTraversal(root.Right)...)
	}
	return result
}
