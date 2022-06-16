package X0094_BinaryTreeInorderTraversal

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
	root := TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	output := inorderTraversal(&root)
	result, _ := json.Marshal(output)
	fmt.Println("Result Binary Tree Inorder Traversal: " + string(result))

}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return []int{}
	}

	if root.Left != nil {
		result = append(result, inorderTraversal(root.Left)...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, inorderTraversal(root.Right)...)
	}
	return result
}
