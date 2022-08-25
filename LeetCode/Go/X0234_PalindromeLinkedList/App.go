package X0234_PalindromeLinkedList

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}

	output := isPalindrome(head)
	fmt.Println("Result Palindrome Linked List: " + strconv.FormatBool(output))

}

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev := &ListNode{-1, nil}

	for slow != nil {
		slowNext := slow.Next
		slow.Next = prev
		prev = slow
		slow = slowNext
	}

	first := head
	second := prev

	for second.Val != -1 {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}
	return true
}
