package X0160_IntersectionOfTwoLinkedLists

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Main() {
	c3 := &ListNode{Val: 3}
	c2 := &ListNode{Val: 2, Next: c3}
	c1 := &ListNode{Val: 1, Next: c2}
	a2 := &ListNode{Val: 2, Next: c1}
	a1 := &ListNode{Val: 1, Next: a2}
	b3 := &ListNode{Val: 3, Next: c1}
	b2 := &ListNode{Val: 2, Next: b3}
	b1 := &ListNode{Val: 1, Next: b2}

	output := getIntersectionNode(a1, b1)
	result, _ := json.Marshal(output)
	fmt.Println("Result Intersection of Two Linked Lists: " + string(result))

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	prevA := []*ListNode{headA}
	prevB := []*ListNode{headB}

	currA := headA
	for currA.Next != nil {
		prevA = append(prevA, currA.Next)
		currA = currA.Next
	}
	currB := headB
	for currB.Next != nil {
		prevB = append(prevB, currB.Next)
		currB = currB.Next
	}

	idxA := len(prevA) - 1
	idxB := len(prevB) - 1
	var intersection *ListNode
	for idxA >= 0 && idxB >= 0 {
		if prevA[idxA] == prevB[idxB] {
			intersection = prevA[idxA]
		}
		idxA--
		idxB--
	}
	return intersection
}
