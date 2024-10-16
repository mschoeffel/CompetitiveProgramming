package X0019_RemoveNthFromEndOfList

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	list     *ListNode
	n        int
	expected *ListNode
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, n: 2,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}, comment: "Example1"},
		{list: &ListNode{Val: 1}, n: 1, expected: nil, comment: "Example2"},
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, n: 1,
			expected: &ListNode{Val: 1}, comment: "Example3"},
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, n: 1,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}, comment: "LastElement"},
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, n: 5,
			expected: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}, comment: "FirstElement"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testApp(element.list, element.n, element.expected, index, t) })
	}
}

func testApp(list *ListNode, n int, expected *ListNode, index int, t *testing.T) {
	actual := removeNthFromEnd(list, n)
	if !reflect.DeepEqual(actual, expected) {
		actualString, _ := json.Marshal(actual)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
