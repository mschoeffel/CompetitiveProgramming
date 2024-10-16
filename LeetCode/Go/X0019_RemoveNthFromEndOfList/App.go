package X0019_RemoveNthFromEndOfList

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
	n := 2

	result, _ := json.Marshal(removeNthFromEnd(&list, n))
	fmt.Println("Result Remove Nth From End of List: " + string(result))

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	_, nextElement := backtrack(head, n)
	return nextElement
}

func backtrack(head *ListNode, n int) (int, *ListNode) {
	backIndex := 0
	nextElement := head.Next
	if head.Next == nil {
		backIndex = 1
	}

	if head.Next != nil {
		backIndex, nextElement = backtrack(head.Next, n)

	}

	head.Next = nextElement

	if backIndex == n {
		nextElement = head.Next
	} else {
		nextElement = head
	}

	return backIndex + 1, nextElement
}
