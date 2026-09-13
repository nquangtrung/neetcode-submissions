/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(depth(root.Right), depth(root.Left)) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := depth(root.Left)
	rightDepth := depth(root.Right)

	maxDiaLeft := diameterOfBinaryTree(root.Left)
	maxDiaRight := diameterOfBinaryTree(root.Right)

	maxChildDia := max(maxDiaLeft, maxDiaRight)
	maxDiaFromDepth := max(leftDepth, rightDepth)


	return max(leftDepth + rightDepth, max(maxChildDia, maxDiaFromDepth))
}
