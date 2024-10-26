package X2458_HeightOfBinaryTreeAfterSubtreeRemovalQueries

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	root     *TreeNode
	queries  []int
	expected []int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 7}}}}, queries: []int{4}, expected: []int{2}, comment: "Example1"},
		{root: &TreeNode{Val: 5, Left: &TreeNode{Val: 8, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}}, queries: []int{3, 2, 4, 8}, expected: []int{3, 2, 3, 2}, comment: "Example2"},
		{root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}}}, queries: []int{1, 4, 3, 4}, expected: []int{2, 1, 1, 1}, comment: "Example3"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.root, element.queries, element.expected, index, t)
			})
	}
}

func testApp(root *TreeNode, queries []int, expected []int, index int, t *testing.T) {
	actual := treeQueries(root, queries)
	if !reflect.DeepEqual(actual, expected) {
		actualStr, _ := json.Marshal(actual)
		expectedStr, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedStr) +
			" Actual: " + string(actualStr))
	}
}
