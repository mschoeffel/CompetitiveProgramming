package X0257_BinaryTreePaths

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	root     *TreeNode
	expected []string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}, expected: []string{"1->2->5", "1->3"}, comment: "Example1"},
		{root: &TreeNode{Val: 1}, expected: []string{"1"}, comment: "Example2"},
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}, expected: []string{"1->2->4", "1->2->5", "1->3"}, comment: "MoreComplex"},
		{root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, expected: []string{"1->2"}, comment: "OneSide"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.root, element.expected, index, t)
			})
	}
}

func testApp(root *TreeNode, expected []string, index int, t *testing.T) {
	result := binaryTreePaths(root)
	if !reflect.DeepEqual(result, expected) {
		actualString, _ := json.Marshal(result)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
