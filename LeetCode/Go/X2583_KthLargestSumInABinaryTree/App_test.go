package X2583_KthLargestSumInABinaryTree

import (
	"strconv"
	"testing"
)

type InputCase struct {
	root     *TreeNode
	k        int
	expected int64
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{root: &TreeNode{Val: 5, Left: &TreeNode{Val: 8, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}}, k: 2, expected: 13, comment: "Example1"},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, k: 1, expected: 3, comment: "Example2"},
		{root: &TreeNode{Val: 5, Left: &TreeNode{Val: 8, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}}, k: 4, expected: -1, comment: "FewerLevels"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.root, element.k, element.expected, index, t)
			})
	}
}

func testApp(root *TreeNode, k int, expected int64, index int, t *testing.T) {
	actual := kthLargestLevelSum(root, k)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(int(expected)) +
			" Actual: " + strconv.Itoa(int(actual)))
	}
}
