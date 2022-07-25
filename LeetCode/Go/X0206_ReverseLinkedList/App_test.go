package X0206_ReverseLinkedList

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	list     *ListNode
	expected *ListNode
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			expected: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}, comment: "Example1"},
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, expected: &ListNode{Val: 2, Next: &ListNode{Val: 1}}, comment: "Example2"},
		{list: nil, expected: nil, comment: "Example3"},
		{list: &ListNode{Val: 1}, expected: &ListNode{Val: 1}, comment: "SingleNode"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.list, element.expected, index, t)
			})
	}

	testSets = []InputCase{
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			expected: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}, comment: "Example1"},
		{list: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, expected: &ListNode{Val: 2, Next: &ListNode{Val: 1}}, comment: "Example2"},
		{list: nil, expected: nil, comment: "Example3"},
		{list: &ListNode{Val: 1}, expected: &ListNode{Val: 1}, comment: "SingleNode"},
	}

	for index, element := range testSets {
		t.Run("Test Recursive: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testAppRecursive(element.list, element.expected, index, t)
			})
	}
}

func testApp(list *ListNode, expected *ListNode, index int, t *testing.T) {
	actual := reverseList(list)
	if !reflect.DeepEqual(actual, expected) {
		resultActual, _ := json.Marshal(actual)
		resultExpected, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(resultExpected) +
			" Actual: " + string(resultActual))
	}
}

func testAppRecursive(list *ListNode, expected *ListNode, index int, t *testing.T) {
	actual := reverseListRecursive(list)
	if !reflect.DeepEqual(actual, expected) {
		resultActual, _ := json.Marshal(actual)
		resultExpected, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(resultExpected) +
			" Actual: " + string(resultActual))
	}
}
