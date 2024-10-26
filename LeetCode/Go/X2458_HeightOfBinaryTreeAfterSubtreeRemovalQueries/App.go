package X2458_HeightOfBinaryTreeAfterSubtreeRemovalQueries

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct {
	maxHeightAfterRemoval [100001]int
	currentMaxHeight      int
}

func Main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 7}}}}
	queries := []int{4}

	result := treeQueries(root, queries)
	resultStr, _ := json.Marshal(result)
	fmt.Println("Result Height of Binary Tree After Subtree Removal Queries: " + string(resultStr))
}

func treeQueries(root *TreeNode, queries []int) []int {
	sol := &Solution{}

	sol.traverseLeftToRight(root, 0)

	sol.currentMaxHeight = 0
	sol.traverseRightToLeft(root, 0)

	queryResults := make([]int, len(queries))
	for i, query := range queries {
		queryResults[i] = sol.maxHeightAfterRemoval[query]
	}

	return queryResults
}

func (s *Solution) traverseLeftToRight(node *TreeNode, currentHeight int) {
	if node == nil {
		return
	}

	s.maxHeightAfterRemoval[node.Val] = s.currentMaxHeight
	s.currentMaxHeight = max(s.currentMaxHeight, currentHeight)

	s.traverseLeftToRight(node.Left, currentHeight+1)
	s.traverseLeftToRight(node.Right, currentHeight+1)
}

func (s *Solution) traverseRightToLeft(node *TreeNode, currentHeight int) {
	if node == nil {
		return
	}

	s.maxHeightAfterRemoval[node.Val] = max(
		s.maxHeightAfterRemoval[node.Val],
		s.currentMaxHeight,
	)

	s.currentMaxHeight = max(currentHeight, s.currentMaxHeight)

	s.traverseRightToLeft(node.Right, currentHeight+1)
	s.traverseRightToLeft(node.Left, currentHeight+1)
}
