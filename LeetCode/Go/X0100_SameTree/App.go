package X0100_SameTree

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
	p := TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	q := TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	result := isSameTree(&p, &q)
	fmt.Println("Result Same Tree: " + strconv.FormatBool(result))

}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
