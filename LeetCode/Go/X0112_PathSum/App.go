package X0112_PathSum

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
	root := TreeNode{Val: 5,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}}}
	targetSum := 22

	result := hasPathSum(&root, targetSum)
	fmt.Println("Result Path Sum: " + strconv.FormatBool(result))

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	target := targetSum - root.Val
	if root.Right == nil && root.Left == nil {
		return target == 0
	}
	return hasPathSum(root.Left, target) || hasPathSum(root.Right, target)
}
