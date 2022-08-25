package X0234_PalindromeLinkedList

import (
	"strconv"
	"testing"
)

type InputCase struct {
	head     *ListNode
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}, expected: true, comment: "Example1"},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, expected: false, comment: "Example2"},
		{head: &ListNode{Val: 1}, expected: true, comment: "SingleNode"},
		{head: nil, expected: true, comment: "Null"},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}}}, expected: true, comment: "MiddleNodeHappyCase"},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}}}}}}, expected: false, comment: "MiddleNodeBadCase"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.head, element.expected, index, t)
			})
	}
}

func testApp(head *ListNode, expected bool, index int, t *testing.T) {
	result := isPalindrome(head)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(result))
	}
}
