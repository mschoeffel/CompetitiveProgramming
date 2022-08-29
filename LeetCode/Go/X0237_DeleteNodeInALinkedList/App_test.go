package X0237_DeleteNodeInALinkedList

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	node     *ListNode
	list     *ListNode
	expected *ListNode
	comment  string
}

func TestApp(t *testing.T) {
	nodeExample1 := &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9}}}
	listExample1 := &ListNode{Val: 4, Next: nodeExample1}

	nodeExample2 := &ListNode{Val: 1, Next: &ListNode{Val: 9}}
	listExample2 := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nodeExample2}}

	nodeTwoNodes := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	listTwoNodes := nodeTwoNodes

	testSets := []InputCase{
		{node: nodeExample1, list: listExample1, expected: &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9}}}, comment: "Example1"},
		{node: nodeExample2, list: listExample2, expected: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}, comment: "Example2"},
		{node: nodeTwoNodes, list: listTwoNodes, expected: &ListNode{Val: 2}, comment: "TwoNodes"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.node, element.list, element.expected, index, t)
			})
	}
}

func testApp(node *ListNode, list *ListNode, expected *ListNode, index int, t *testing.T) {
	deleteNode(node)
	if !reflect.DeepEqual(list, expected) {
		actualString, _ := json.Marshal(list)
		expectedString, _ := json.Marshal(expected)
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + string(expectedString) +
			" Actual: " + string(actualString))
	}
}
