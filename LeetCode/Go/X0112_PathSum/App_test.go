package X0112_PathSum

import (
	"strconv"
	"testing"
)

type InputCase struct {
	root     *TreeNode
	target   int
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{root: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
			Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}}}, target: 22, expected: true, comment: "Example1"},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, target: 5, expected: false, comment: "Example2"},
		{root: &TreeNode{Val: 4}, target: 4, expected: true, comment: "SingleRoot"},
		{root: nil, target: 0, expected: false, comment: "NilRoot"},
		{root: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
			Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: -2}, Right: &TreeNode{Val: 1}}}}, target: 15, expected: true, comment: "NegativeVals"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.root, element.target, element.expected, index, t)
			})
	}
}

func testApp(root *TreeNode, target int, expected bool, index int, t *testing.T) {
	actual := hasPathSum(root, target)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
