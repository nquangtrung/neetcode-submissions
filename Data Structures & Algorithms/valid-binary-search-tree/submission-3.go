/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func checkValidBranch(root *TreeNode, minValue int, maxValue int) bool {
	if root == nil {
		return true
	}

	if root.Val <= minValue || root.Val >= maxValue {
		return false
	}

	return checkValidBranch(root.Left, minValue, root.Val) && checkValidBranch(root.Right, root.Val, maxValue)
}

func isValidBST(root *TreeNode) bool {
	return checkValidBranch(root, math.MinInt, math.MaxInt)
}
