package X0951_FlipEquivalentBinaryTrees

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
	root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 8}}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}}
	root2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 7}, Left: &TreeNode{Val: 8}}}, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}}

	result := flipEquiv(root1, root2)
	fmt.Println("Result Flip Equivalent Binary Trees: " + strconv.FormatBool(result))
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 != nil && root2 != nil {
		if root1.Val == root2.Val {
			if flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) {
				return true
			} else if flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left) {
				return true
			}
		}
	}
	return false
}
