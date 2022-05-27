package X0021_MergeTwoSortedLists

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Main() {
	listOne := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	listTwo := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	result, _ := json.Marshal(mergeTwoLists(&listOne, &listTwo))
	fmt.Println("Result Merge Two Sorted Lists: " + string(result))

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	listOnePointer := list1
	listTwoPointer := list2

	resultHead := &ListNode{}
	resultPointer := resultHead
	if listOnePointer != nil {
		for listOnePointer != nil {
			if listTwoPointer == nil || listOnePointer.Val <= listTwoPointer.Val {
				resultPointer.Next = &ListNode{Val: listOnePointer.Val}
				listOnePointer = listOnePointer.Next
			} else {
				resultPointer.Next = &ListNode{Val: listTwoPointer.Val}
				listTwoPointer = listTwoPointer.Next
			}
			resultPointer = resultPointer.Next
		}
		resultPointer.Next = listTwoPointer
	} else {
		resultPointer.Next = listTwoPointer
	}
	return resultHead.Next
}
