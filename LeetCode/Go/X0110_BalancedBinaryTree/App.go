package X0110_BalancedBinaryTree

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
	root := TreeNode{Val: 3, Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}

	result := isBalanced(&root)
	fmt.Println("Result Balanced Binary Tree: " + strconv.FormatBool(result))

}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	validLeft, depthLeft := countDepth(root.Left)
	validRight, depthRight := countDepth(root.Right)
	diff := depthLeft - depthRight
	return validLeft && validRight && diff <= 1 && diff >= -1
}

func countDepth(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	depthRight := 0
	validRight := true
	if root.Right != nil {
		validRight, depthRight = countDepth(root.Right)
	}
	depthLeft := 0
	validLeft := true
	if root.Left != nil {
		validLeft, depthLeft = countDepth(root.Left)
	}
	diff := depthLeft - depthRight
	valid := validRight && validLeft && diff <= 1 && diff >= -1

	if depthLeft > depthRight {
		return valid, depthLeft + 1
	} else {
		return valid, depthRight + 1
	}
}
