package X0404_SumOfLeftLeaves

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
		{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, expected: 24, comment: "Example1"},
		{root: &TreeNode{Val: 1}, expected: 0, comment: "Example2"},
		{root: nil, expected: 0, comment: "NilRoot"},
		{root: &TreeNode{Val: 3, Right: &TreeNode{Val: 20, Right: &TreeNode{Val: 7}}}, expected: 0, comment: "NoneRightLeave"},
		{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 10}}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, expected: 25, comment: "LeftWing"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.root, element.expected, index, t)
			})
	}
}

func testApp(root *TreeNode, expected int, index int, t *testing.T) {
	actual := sumOfLeftLeaves(root)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.Itoa(expected) +
			" Actual: " + strconv.Itoa(actual))
	}
}
