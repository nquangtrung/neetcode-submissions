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
	if p.Val > q.Val {
		return lowestCommonAncestor(root, q, p)
	}

	current := root
	for current != nil {
		if p.Val <= current.Val && q.Val >= current.Val {
			// This is the lowest common ancestor
			return current
		} else if p.Val <= current.Val && q.Val <= current.Val {
			current = current.Left
		} else if p.Val >= current.Val && q.Val >= current.Val {
			current = current.Right
		}
	}

	return nil
}
