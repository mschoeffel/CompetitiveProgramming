package X0404_SumOfLeftLeaves

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
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 20}, Right: &TreeNode{Val: 7}}}

	result := sumOfLeftLeaves(root)
	fmt.Println("Result Sum of Left Leaves: " + strconv.Itoa(result))

}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	} else {
		sum += sumOfLeftLeaves(root.Left)
	}
	sum += sumOfLeftLeaves(root.Right)

	return sum
}
