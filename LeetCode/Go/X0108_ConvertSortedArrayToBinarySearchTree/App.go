package X0108_ConvertSortedArrayToBinarySearchTree

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Main() {
	nums := []int{-10, -3, 0, 5, 9}

	output := sortedArrayToBST(nums)
	result, _ := json.Marshal(output)
	fmt.Println("Result Convert Sorted Array to Binary Tree: " + string(result))

}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	center := len(nums) / 2
	leftSide := nums[:center]
	rightSide := nums[center+1:]
	return &TreeNode{Val: nums[center], Left: sortedArrayToBST(leftSide), Right: sortedArrayToBST(rightSide)}
}
