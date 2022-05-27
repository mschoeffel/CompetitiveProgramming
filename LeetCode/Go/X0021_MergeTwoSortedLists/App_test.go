package X0021_MergeTwoSortedLists

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	listOne  *ListNode
	listTwo  *ListNode
	expected *ListNode
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{listOne: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
			listTwo:  &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}}, comment: "Example1"},
		{listOne: nil,
			listTwo:  nil,
			expected: nil, comment: "Example2"},
		{listOne: nil,
			listTwo:  &ListNode{Val: 0},
			expected: &ListNode{Val: 0}, comment: "Example3"},
		{listOne: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
			listTwo:  &ListNode{Val: 3, Next: &ListNode{Val: 5}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, comment: "ListOneLonger"},
		{listOne: &ListNode{Val: 3, Next: &ListNode{Val: 5}},
			listTwo:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, comment: "ListTwoLonger"},
		{listOne: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}},
			listTwo:  &ListNode{Val: 3, Next: &ListNode{Val: 5}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}, comment: "ListOneLongerAndEnding"},
		{listOne: &ListNode{Val: 3, Next: &ListNode{Val: 5}},
			listTwo:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}, comment: "ListTwoLongerAndEnding"},
		{listOne: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			listTwo:  &ListNode{Val: 3, Next: &ListNode{Val: 4}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}, comment: "ListOneStarting"},
		{listOne: &ListNode{Val: 3, Next: &ListNode{Val: 4}},
			listTwo:  &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}, comment: "ListTwoStarting"},
		{listOne: nil,
			listTwo:  &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, comment: "ListOneEmpty"},
		{listOne: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			listTwo:  nil,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, comment: "ListTwoEmpty"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testApp(element.listOne, element.listTwo, element.expected, index, t) })
	}
}

func testApp(listOne *ListNode, listTwo *ListNode, expected *ListNode, index int, t *testing.T) {
	actual := mergeTwoLists(listOne, listTwo)
	if !reflect.DeepEqual(actual, expected) {
		actualString, _ := json.Marshal(actual)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
