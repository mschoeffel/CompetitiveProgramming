package X0237_DeleteNodeInALinkedList

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Main() {
	node := &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9}}}
	list := &ListNode{Val: 4, Next: node}

	deleteNode(node)
	result, _ := json.Marshal(list)
	fmt.Println("Result Delete Node in a Linked List: " + string(result))

}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
