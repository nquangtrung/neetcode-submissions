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

	return max(depth(root.Left), depth(root.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	
    leftDepth := depth(root.Left)
	rightDepth := depth(root.Right)

	// fmt.Printf("left %d, right %d\n", leftDepth, rightDepth)

	return math.Abs(float64(leftDepth) - float64(rightDepth)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
