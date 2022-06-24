package X0111_MinimumDepthOfBinaryTree

import (
	"strconv"
	"testing"
)

type InputCase struct {
	root     *TreeNode
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, expected: 2, comment: "Example1"},
		{root: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 5}}}}}, expected: 5, comment: "Example2"},
		{root: &TreeNode{Val: 2}, expected: 1, comment: "SingleNode"},
		{root: nil, expected: 0, comment: "NilRoot"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.root, element.expected, index, t)
			})
	}
}

func testApp(root *TreeNode, expected int, index int, t *testing.T) {
	actual := minDepth(root)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
