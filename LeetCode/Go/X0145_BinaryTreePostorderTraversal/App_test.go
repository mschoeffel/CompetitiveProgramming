package X0145_BinaryTreePostorderTraversal

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	root     *TreeNode
	expected []int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{root: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, expected: []int{3, 2, 1}, comment: "Example1"},
		{root: nil, expected: []int{}, comment: "Example2"},
		{root: &TreeNode{Val: 1}, expected: []int{1}, comment: "Example3"},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5, Right: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 6}},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 9}}}}, expected: []int{7, 5, 6, 4, 3, 9, 8, 2, 1}, comment: "MoreComplex"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.root, element.expected, index, t)
			})
	}
}

func testApp(root *TreeNode, expected []int, index int, t *testing.T) {
	result := postorderTraversal(root)
	if !reflect.DeepEqual(result, expected) {
		actualString, _ := json.Marshal(result)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
