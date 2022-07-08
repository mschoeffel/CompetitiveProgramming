package X0145_BinaryTreePostorderTraversal

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

	output := postorderTraversal(root)
	result, _ := json.Marshal(output)
	fmt.Println("Result Binary Tree Postorder Traversal: " + string(result))

}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	result = append(result, postorderTraversal(root.Left)...)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}
