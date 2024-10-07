package X0002_AddTwoNumbers

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Main() {
	example1l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	example1l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	result, _ := json.Marshal(addTwoNumbers(example1l1, example1l2))
	fmt.Println("Result Add Two Numbers: " + string(result))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	var result []int

	for l1 != nil || l2 != nil {
		l1Val := 0
		l2Val := 0
		if l1 != nil {
			l1Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val += l2.Val
			l2 = l2.Next
		}

		res := l1Val + l2Val + carry
		if res >= 10 {
			res -= 10
			carry = 1
		} else {
			carry = 0
		}
		result = append(result, res)
	}

	if carry == 1 {
		result = append(result, 1)
	}

	return createListNode(result)
}

func createListNode(arr []int) *ListNode {
	head := &ListNode{}
	current := head
	for _, val := range arr {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head.Next
}
