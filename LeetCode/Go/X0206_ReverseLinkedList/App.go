package X0206_ReverseLinkedList

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Main() {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	output := reverseList(list)
	result, _ := json.Marshal(output)
	fmt.Println("Result Reverse Linked List: " + string(result))

	list = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	output = reverseList(list)
	result, _ = json.Marshal(output)
	fmt.Println("Result Reverse Linked List Recursive: " + string(result))
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	currentNode := head
	var previousNode *ListNode
	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	return previousNode
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	_, start := reverseListRecursiveSub(head)
	return start
}

func reverseListRecursiveSub(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}
	node, start := reverseListRecursiveSub(head.Next)
	node.Next = head
	head.Next = nil
	return node.Next, start
}
