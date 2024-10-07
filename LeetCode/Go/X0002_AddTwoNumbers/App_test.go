package X0002_AddTwoNumbers

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	l1       *ListNode
	l2       *ListNode
	expected *ListNode
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}, expected: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}, comment: "Example1"},
		{l1: &ListNode{Val: 0}, l2: &ListNode{Val: 0}, expected: &ListNode{Val: 0}, comment: "Example2"},
		{l1: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}, l2: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}, expected: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}}}}, comment: "Example3"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.l1, element.l2, element.expected, index, t)
			})
	}
}

func testApp(l1 *ListNode, l2 *ListNode, expected *ListNode, index int, t *testing.T) {
	actual := addTwoNumbers(l1, l2)
	if !reflect.DeepEqual(actual, expected) {
		actualString, _ := json.Marshal(actual)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
