package X0110_BalancedBinaryTree

import (
	"strconv"
	"testing"
)

type InputCase struct {
	root     *TreeNode
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, expected: true, comment: "Example1"},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2}}, expected: false, comment: "Example2"},
		{root: nil, expected: true, comment: "NullRoot"},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5}}}}, expected: false, comment: "SimpleBadCase"},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 8}, Right: &TreeNode{Val: 12, Left: &TreeNode{Val: 10}}},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5}}}}, expected: true, comment: "SimpleHappyCase"},
		{root: &TreeNode{Val: 1}, expected: true, comment: "SingleNode"},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}}, expected: false, comment: "SingleNode"},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 8}, Right: &TreeNode{Val: 12, Left: &TreeNode{Val: 10}}},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}}}, expected: false, comment: "DeepInvalid"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.root, element.expected, index, t)
			})
	}
}

func testApp(root *TreeNode, expected bool, index int, t *testing.T) {
	actual := isBalanced(root)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
