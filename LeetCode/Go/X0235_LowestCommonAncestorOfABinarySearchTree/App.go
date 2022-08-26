package X0235_LowestCommonAncestorOfABinarySearchTree

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
	root := &TreeNode{Val: 1}
	p := &TreeNode{Val: 2}
	q := &TreeNode{Val: 8}

	output := lowestCommonAncestor(root, p, q)
	fmt.Println("Result Lowest Common Ancestor of a Binary Search Tree: " + strconv.Itoa(output.Val))

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if (p.Val >= root.Val && q.Val <= root.Val) || (q.Val >= root.Val && p.Val <= root.Val) {
		return root
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}
