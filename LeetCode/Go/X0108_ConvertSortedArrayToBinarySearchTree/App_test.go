package X0108_ConvertSortedArrayToBinarySearchTree

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	nums     []int
	expected *TreeNode
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{nums: []int{-10, -3, 0, 5, 9}, expected: &TreeNode{Val: 0, Left: &TreeNode{Val: -3, Left: &TreeNode{Val: -10}}, Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}}}, comment: "Example1"},
		{nums: []int{1, 3}, expected: &TreeNode{Val: 3, Left: &TreeNode{Val: 1}}, comment: "Example2"},
		{nums: []int{1}, expected: &TreeNode{Val: 1}, comment: "SingleNode"},
		{nums: []int{-10, -3, -1, 0, 5, 9}, expected: &TreeNode{Val: 0, Left: &TreeNode{Val: -3, Left: &TreeNode{Val: -10}, Right: &TreeNode{Val: -1}}, Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}}}, comment: "EvenLength"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.nums, element.expected, index, t)
			})
	}
}

func testApp(nums []int, expected *TreeNode, index int, t *testing.T) {
	actual := sortedArrayToBST(nums)
	if !reflect.DeepEqual(actual, expected) {
		resultActual, _ := json.Marshal(actual)
		resultExpected, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(resultExpected) +
			" Actual: " + string(resultActual))
	}
}
