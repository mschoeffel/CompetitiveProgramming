package X0257_BinaryTreePaths

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Main() {
	root := &TreeNode{Val: 1}

	output := binaryTreePaths(root)
	result, _ := json.Marshal(output)
	fmt.Println("Result Binary Tree Paths: " + string(result))

}

func binaryTreePaths(root *TreeNode) []string {
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	var paths []string
	if root.Left != nil {
		paths = append(paths, binaryTreePathsSumUp(root.Left, strconv.Itoa(root.Val))...)
	}
	if root.Right != nil {
		paths = append(paths, binaryTreePathsSumUp(root.Right, strconv.Itoa(root.Val))...)
	}
	return paths
}

func binaryTreePathsSumUp(node *TreeNode, current string) []string {
	if node.Left == nil && node.Right == nil {
		return []string{current + "->" + strconv.Itoa(node.Val)}
	}

	var paths []string
	if node.Left != nil {
		paths = append(paths, binaryTreePathsSumUp(node.Left, current+"->"+strconv.Itoa(node.Val))...)
	}
	if node.Right != nil {
		paths = append(paths, binaryTreePathsSumUp(node.Right, current+"->"+strconv.Itoa(node.Val))...)
	}

	return paths
}
