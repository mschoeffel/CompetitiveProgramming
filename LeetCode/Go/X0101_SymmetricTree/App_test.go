package X0101_SymmetricTree

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
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}, expected: true, comment: "Example1"},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
			Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}, expected: false, comment: "Example2"},
		{root: &TreeNode{Val: 1}, expected: true, comment: "EmptyRoot"},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2},
			Right: &TreeNode{Val: 2}}, expected: true, comment: "SimpleTree"},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 7}}},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Right: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 3}}}, expected: true, comment: "MoreComplexHappyCase"},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 7}}},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Right: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 4}}}, expected: false, comment: "MoreComplexBadCase"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.root, element.expected, index, t)
			})
	}
}

func testApp(root *TreeNode, expected bool, index int, t *testing.T) {
	actual := isSymmetric(root)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
