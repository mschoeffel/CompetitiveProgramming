package X0111_MinimumDepthOfBinaryTree

import (
	"fmt"
	"math"
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

	result := minDepth(&root)
	fmt.Println("Result Minimum Depth of Binary Tree: " + strconv.Itoa(result))

}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	depthLeft := math.MaxInt32
	if root.Left != nil {
		depthLeft = minDepth(root.Left)
	}
	depthRight := math.MaxInt32
	if root.Right != nil {
		depthRight = minDepth(root.Right)
	}
	if depthRight > depthLeft {
		return depthLeft + 1
	} else {
		return depthRight + 1
	}
}
