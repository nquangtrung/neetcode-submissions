/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
   	if p == nil && q == nil {
		return true
   	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if !isSameTree(p.Left, q.Left) {
		return false
	}
	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
   	if root == nil && subRoot == nil {
		return true
   	}
	if root == nil && subRoot != nil {
		return false
	}
	if root != nil && subRoot == nil {
		return true
	}
	if isSameTree(root, subRoot) {
		return true
	}
	if isSameTree(root.Left, subRoot) {
		// fmt.Printf("%d: left is same tree as sub", root.Val)
		return true
	}
	if isSameTree(root.Right, subRoot) {
		// fmt.Printf("%d: right is same tree as sub", root.Val)
		return true
	}
	if isSubtree(root.Left, subRoot) {
		// fmt.Printf("%d: left contains sub", root.Val)
		return true
	}
	if isSubtree(root.Right, subRoot) {
		// fmt.Printf("%d: right contains sub", root.Val)
		return true
	}

	return false
}
