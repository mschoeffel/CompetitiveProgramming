package X0203_RemoveLinkedListElements

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Main() {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}

	output := removeElements(list, 6)
	result, _ := json.Marshal(output)
	fmt.Println("Result Remove Linked List Elements: " + string(result))
}

func removeElements(head *ListNode, val int) *ListNode {
	currentElement := head
	var start *ListNode = nil
	var pointer *ListNode = nil
	for currentElement != nil {
		if currentElement.Val != val {
			if start == nil {
				start = currentElement
			}
			if pointer == nil {
				pointer = currentElement
			} else {
				pointer.Next = currentElement
				pointer = currentElement
			}
		}
		currentElement = currentElement.Next
	}
	if pointer != nil {
		pointer.Next = nil
	}
	return start
}
