package X0104_MaximumDepthOfBinaryTree

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Main() {
	root := TreeNode{Val: 1, Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}

	result := maxDepth(&root)
	fmt.Println("Result Maximum Depth of Binary Tree: " + strconv.Itoa(result))

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depthLeft := maxDepth(root.Left)
	depthRight := maxDepth(root.Right)
	if depthRight > depthLeft {
		return depthRight + 1
	} else {
		return depthLeft + 1
	}
}
