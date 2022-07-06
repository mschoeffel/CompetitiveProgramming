package X0141_LinkedListCycle

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Main() {
	first := &ListNode{Val: 3}
	second := &ListNode{Val: 2}
	third := &ListNode{Val: 0}
	fourth := &ListNode{Val: -4, Next: second}
	third.Next = fourth
	second.Next = third
	first.Next = second

	result := hasCycle(first)
	fmt.Println("Result Linked List Cycle: " + strconv.FormatBool(result))
	result = hasCycleFastSlow(first)
	fmt.Println("Result Linked List Cycle Fast Slow: " + strconv.FormatBool(result))

}

func hasCycle(head *ListNode) bool {
	return hasCycleNode(head, map[*ListNode]bool{})
}

func hasCycleNode(node *ListNode, visited map[*ListNode]bool) bool {
	if node == nil {
		return false
	}
	if _, ok := visited[node]; ok {
		return true
	} else {
		visited[node] = true
		if node.Next != nil {
			return hasCycleNode(node.Next, visited)
		} else {
			return false
		}
	}
}

func hasCycleFastSlow(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}
