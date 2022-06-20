package X0101_SymmetricTree

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
	root := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}

	result := isSymmetric(&root)
	fmt.Println("Result Symmetric Tree: " + strconv.FormatBool(result))

}

func isSymmetric(root *TreeNode) bool {
	return isNodeSymmetric(root.Right, root.Left)
}

func isNodeSymmetric(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isNodeSymmetric(p.Right, q.Left) && isNodeSymmetric(p.Left, q.Right)
}
