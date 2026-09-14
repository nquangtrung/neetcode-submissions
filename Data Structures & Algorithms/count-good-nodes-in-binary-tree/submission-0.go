/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func countGoodNodes(root *TreeNode, maxValue int) int {
	if root == nil {
		return 0
	}

	result := 0
	if root.Val >= maxValue {
		result += 1
	}

	return result + countGoodNodes(root.Left, max(root.Val, maxValue)) + countGoodNodes(root.Right, max(root.Val, maxValue))
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countGoodNodes(root, root.Val) 
}
