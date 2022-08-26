package X0235_LowestCommonAncestorOfABinarySearchTree

import (
	"strconv"
	"testing"
)

type InputCase struct {
	root     *TreeNode
	p        *TreeNode
	q        *TreeNode
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{root: &TreeNode{Val: 6, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}}, p: &TreeNode{Val: 2}, q: &TreeNode{Val: 8}, expected: 6, comment: "Example1"},
		{root: &TreeNode{Val: 6, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}}, p: &TreeNode{Val: 2}, q: &TreeNode{Val: 4}, expected: 2, comment: "Example2"},
		{root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, p: &TreeNode{Val: 2}, q: &TreeNode{Val: 1}, expected: 2, comment: "Example3"},
		{root: &TreeNode{Val: 6, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}}, p: &TreeNode{Val: 0}, q: &TreeNode{Val: 5}, expected: 2, comment: "MultiLayer"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.root, element.p, element.q, element.expected, index, t)
			})
	}
}

func testApp(root *TreeNode, p *TreeNode, q *TreeNode, expected int, index int, t *testing.T) {
	result := lowestCommonAncestor(root, p, q)
	if result.Val != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(result.Val))
	}
}
