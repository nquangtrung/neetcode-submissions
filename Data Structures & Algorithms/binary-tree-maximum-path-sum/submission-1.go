/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func calculate(root *TreeNode) (int, int) {
	if root == nil {
		return math.MinInt, math.MinInt
	}

	leftDeepPathSum, maxLeftPathSum := calculate(root.Left)
	rightDeepPathSum, maxRightPathSum := calculate(root.Right)

	deepPathSum := root.Val
	if leftDeepPathSum > 0 || rightDeepPathSum > 0 {
		deepPathSum += max(leftDeepPathSum, rightDeepPathSum)
	}

	maxPathSum := max(maxLeftPathSum, maxRightPathSum)
	newPathSum := root.Val
	if leftDeepPathSum > 0 {
		newPathSum += leftDeepPathSum
	}
	if rightDeepPathSum > 0 {
		newPathSum += rightDeepPathSum
	}

	return deepPathSum, max(maxPathSum, newPathSum)
}

func maxPathSum(root *TreeNode) int {
	_, result := calculate(root)
	return result
}
