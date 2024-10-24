package X0951_FlipEquivalentBinaryTrees

import (
	"strconv"
	"testing"
)

type InputCase struct {
	root1    *TreeNode
	root2    *TreeNode
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{root1: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 8}}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}}, root2: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 7}, Left: &TreeNode{Val: 8}}}, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}}, expected: true, comment: "Example1"},
		{root1: nil, root2: nil, expected: true, comment: "Example2"},
		{root1: &TreeNode{Val: 1}, root2: nil, expected: false, comment: "Example3"},
		{root1: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 8}}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}}, root2: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 7}, Left: &TreeNode{Val: 8}}}, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}}}, expected: false, comment: "WrongFurtherDown"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.root1, element.root2, element.expected, index, t)
			})
	}
}

func testApp(root1 *TreeNode, root2 *TreeNode, expected bool, index int, t *testing.T) {
	actual := flipEquiv(root1, root2)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
