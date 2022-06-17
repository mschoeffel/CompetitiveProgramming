package X0100_SameTree

import (
	"strconv"
	"testing"
)

type InputCase struct {
	p        *TreeNode
	q        *TreeNode
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			q: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, expected: true, comment: "Example1"},
		{p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			q: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, expected: false, comment: "Example2"},
		{p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
			q: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, expected: false, comment: "Example3"},
		{p: &TreeNode{}, q: &TreeNode{}, expected: true, comment: "BothEmpty"},
		{p: &TreeNode{}, q: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, expected: false, comment: "PEmpty"},
		{p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, q: &TreeNode{}, expected: false, comment: "QEmpty"},
		{p: nil, q: nil, expected: true, comment: "BothNil"},
		{p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, q: nil, expected: false, comment: "PNil"},
		{p: nil, q: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, expected: false, comment: "QNil"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.p, element.q, element.expected, index, t)
			})
	}
}

func testApp(p *TreeNode, q *TreeNode, expected bool, index int, t *testing.T) {
	actual := isSameTree(p, q)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(actual))
	}
}
