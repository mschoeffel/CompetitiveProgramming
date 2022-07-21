package X0203_RemoveLinkedListElements

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	list     *ListNode
	val      int
	expected *ListNode
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}, val: 6,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, comment: "Example1"},
		{list: nil, val: 1, expected: nil, comment: "Example2"},
		{list: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7}}}}, val: 7,
			expected: nil, comment: "Example3"},
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}}}}, val: 1,
			expected: &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}, comment: "BeginningAndEnd"},
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}, val: 8,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}, comment: "NoRemove"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.list, element.val, element.expected, index, t)
			})
	}
}

func testApp(list *ListNode, val int, expected *ListNode, index int, t *testing.T) {
	actual := removeElements(list, val)
	if !reflect.DeepEqual(actual, expected) {
		resultActual, _ := json.Marshal(actual)
		resultExpected, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(resultExpected) +
			" Actual: " + string(resultActual))
	}
}
