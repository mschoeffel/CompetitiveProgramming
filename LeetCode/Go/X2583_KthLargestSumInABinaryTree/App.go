package X2583_KthLargestSumInABinaryTree

import (
	"fmt"
	"slices"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Main() {
	root := &TreeNode{Val: 5, Left: &TreeNode{Val: 8, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}}

	result := kthLargestLevelSum(root, 2)
	fmt.Println("Result Kth Largest Sum in a Binary Tree: " + strconv.Itoa(int(result)))
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return 0
	}

	levelSum := make(map[int]int64)
	nextNodes := make([]*TreeNode, 0)
	currentNodes := make([]*TreeNode, 0)
	currentNodes = append(currentNodes, root)
	idx := 0

	for len(currentNodes) > 0 {
		for _, node := range currentNodes {
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}
			levelSum[idx] += int64(node.Val)
		}
		idx++
		currentNodes = nextNodes
		nextNodes = make([]*TreeNode, 0)
	}

	sumArray := make([]int64, 0)
	for _, sum := range levelSum {
		sumArray = append(sumArray, sum)
	}

	if k > len(sumArray) {
		return -1
	}

	slices.Sort(sumArray)
	return sumArray[len(sumArray)-k]
}
