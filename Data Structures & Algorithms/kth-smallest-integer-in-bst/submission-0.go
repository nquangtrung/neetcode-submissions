/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	if root.Left != nil {
		result = append(result, traverse(root.Left)...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, traverse(root.Right)...)
	}

	return result
}

func kthSmallest(root *TreeNode, k int) int {
	arr := traverse(root)

	return arr[k - 1]
}