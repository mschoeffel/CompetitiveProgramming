package X0226_InvertBinaryTree

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
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}

	output := invertTree(root)
	result, _ := json.Marshal(output)
	fmt.Println("Result Invert Binary Tree: " + string(result))

}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	temp := invertTree(root.Left)
	root.Left = invertTree(root.Right)
	root.Right = temp
	return root
}
