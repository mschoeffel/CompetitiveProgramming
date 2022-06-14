package X0083_RemoveDuplicatesFromSortedList

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Main() {
	list := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}

	result, _ := json.Marshal(deleteDuplicates(&list))
	fmt.Println("Result Remove Duplicates from Sorted List: " + string(result))

}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	currentNode := head.Next
	pointer := &ListNode{Val: head.Val}
	pointerHead := pointer
	for currentNode != nil {
		if pointer.Val != currentNode.Val {
			pointer.Next = &ListNode{Val: currentNode.Val}
			pointer = pointer.Next
		}
		currentNode = currentNode.Next
	}
	return pointerHead
}

func deleteDuplicatesInPlace(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	currentNode := head.Next
	pointer := head
	pointerHead := pointer
	for currentNode != nil {
		if pointer.Val != currentNode.Val {
			pointer.Next = currentNode
			pointer = pointer.Next
		}
		currentNode = currentNode.Next
	}
	pointer.Next = nil
	return pointerHead
}
