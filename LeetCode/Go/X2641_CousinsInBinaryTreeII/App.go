package X2641_CousinsInBinaryTreeII

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
	root := &TreeNode{Val: 5, Left: &TreeNode{Val: 8, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}}

	result, _ := json.Marshal(replaceValueInTree(root))
	fmt.Println("Result Cousins in Binary Tree II: " + string(result))
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	root.Val = 0

	for len(queue) > 0 {
		lengthQueue := len(queue)
		sum := 0

		for i := range lengthQueue {
			currentNode := queue[i]
			if currentNode.Left != nil {
				sum += currentNode.Left.Val
			}
			if currentNode.Right != nil {
				sum += currentNode.Right.Val
			}
		}

		for i := range lengthQueue {
			currentNode := queue[i]
			leftValue, rightValue := 0, 0
			if currentNode.Left != nil {
				leftValue = currentNode.Left.Val
			}
			if currentNode.Right != nil {
				rightValue = currentNode.Right.Val
			}
			if currentNode.Left != nil {
				currentNode.Left.Val = sum - leftValue - rightValue
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				currentNode.Right.Val = sum - leftValue - rightValue
				queue = append(queue, currentNode.Right)
			}
		}
		queue = queue[lengthQueue:]
	}
	return root
}
