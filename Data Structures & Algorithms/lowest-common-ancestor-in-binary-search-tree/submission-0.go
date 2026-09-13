/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isAncestor(root *TreeNode, p *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == p {
		return true
	}

	return isAncestor(root.Left, p) || isAncestor(root.Right, p)
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	
	leftLowestAncestor := lowestCommonAncestor(root.Left, p, q)
	if leftLowestAncestor  != nil {
		return leftLowestAncestor	
	}

	rightLowestAncestor := lowestCommonAncestor(root.Right, p, q)
	if rightLowestAncestor != nil {
		return rightLowestAncestor
	}

	if isAncestor(root, p) && isAncestor(root, q) {
		return root
	}

	return nil
}
