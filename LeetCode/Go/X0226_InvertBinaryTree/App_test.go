package X0226_InvertBinaryTree

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	root     *TreeNode
	expected *TreeNode
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{root: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}},
			expected: &TreeNode{Val: 4, Left: &TreeNode{Val: 7, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}}, comment: "Example1"},
		{root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			expected: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}, comment: "Example2"},
		{root: nil, expected: nil, comment: "Example3"},
		{root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}, expected: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 1}}, comment: "UnevenTree"},
		{root: &TreeNode{Val: 3}, expected: &TreeNode{Val: 3}, comment: "OnlyRoot"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.root, element.expected, index, t)
			})
	}
}

func testApp(root *TreeNode, expected *TreeNode, index int, t *testing.T) {
	result := invertTree(root)
	if !reflect.DeepEqual(result, expected) {
		actualString, _ := json.Marshal(result)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
