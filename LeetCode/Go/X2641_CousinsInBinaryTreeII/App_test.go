package X2641_CousinsInBinaryTreeII

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
		{root: &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 10}}, Right: &TreeNode{Val: 9, Right: &TreeNode{Val: 7}}}, expected: &TreeNode{Val: 0, Left: &TreeNode{Val: 0, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 0, Right: &TreeNode{Val: 11}}}, comment: "Example1"},
		{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, expected: &TreeNode{Val: 0, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 0}}, comment: "Example2"},
		{root: &TreeNode{Val: 3}, expected: &TreeNode{Val: 0}, comment: "SingleRoot"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.root, element.expected, index, t)
			})
	}
}

func testApp(root *TreeNode, expected *TreeNode, index int, t *testing.T) {
	actual := replaceValueInTree(root)
	actualStr, _ := json.Marshal(actual)
	expectedStr, _ := json.Marshal(expected)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedStr) +
			" Actual: " + string(actualStr))
	}
}
