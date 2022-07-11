package X0160_IntersectionOfTwoLinkedLists

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	a        *ListNode
	b        *ListNode
	expected *ListNode
	comment  string
}

func TestApp(t *testing.T) {
	e1c3 := &ListNode{Val: 3}
	e1c2 := &ListNode{Val: 2, Next: e1c3}
	e1c1 := &ListNode{Val: 1, Next: e1c2}
	e1a2 := &ListNode{Val: 2, Next: e1c1}
	e1a1 := &ListNode{Val: 1, Next: e1a2}
	e1b3 := &ListNode{Val: 3, Next: e1c1}
	e1b2 := &ListNode{Val: 2, Next: e1b3}
	e1b1 := &ListNode{Val: 1, Next: e1b2}

	e2c2 := &ListNode{Val: 4}
	e2c1 := &ListNode{Val: 2, Next: e2c2}
	e2a3 := &ListNode{Val: 2, Next: e2c1}
	e2a2 := &ListNode{Val: 2, Next: e2a3}
	e2a1 := &ListNode{Val: 2, Next: e2a2}
	e2b1 := &ListNode{Val: 3, Next: e2c1}

	e3a3 := &ListNode{Val: 4}
	e3a2 := &ListNode{Val: 6, Next: e3a3}
	e3a1 := &ListNode{Val: 2, Next: e3a2}
	e3b2 := &ListNode{Val: 5}
	e3b1 := &ListNode{Val: 1, Next: e3b2}

	e5c1 := &ListNode{Val: 4}

	testSets := []InputCase{
		{a: e1a1, b: e1b1, expected: e1c1, comment: "Example1"},
		{a: e2a1, b: e2b1, expected: e2c1, comment: "Example2"},
		{a: e3a1, b: e3b1, expected: nil, comment: "Example3"},
		{a: nil, b: nil, expected: nil, comment: "NilCase"},
		{a: e5c1, b: e5c1, expected: e5c1, comment: "SingleElement"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.a, element.b, element.expected, index, t)
			})
	}
}

func testApp(a *ListNode, b *ListNode, expected *ListNode, index int, t *testing.T) {
	result := getIntersectionNode(a, b)
	if !reflect.DeepEqual(result, expected) {
		actualString, _ := json.Marshal(result)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
