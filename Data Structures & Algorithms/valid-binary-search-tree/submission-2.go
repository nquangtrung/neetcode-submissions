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

	// fmt.Printf("checking for branch %d:%d:%d\n", root.Val, minValue, maxValue)
	if root.Left != nil  {
		// fmt.Printf("left: %d\n", root.Left.Val)	
		if root.Left.Val >= root.Val || root.Left.Val <= minValue {
			// fmt.Printf("left is valid balance: %d\n", root.Left.Val)
			return false
		}
	}
	
	if root.Right != nil {
		// fmt.Printf("right: %d\n", root.Right.Val)	
		if root.Right.Val <= root.Val || root.Right.Val >= maxValue {
			// fmt.Printf("right is not balance: %d\n", root.Right.Val)
			return false
		}
	}

	return checkValidBranch(root.Left, minValue, root.Val) && checkValidBranch(root.Right, root.Val, maxValue)
}

func isValidBST(root *TreeNode) bool {
	return checkValidBranch(root, math.MinInt, math.MaxInt)
}
