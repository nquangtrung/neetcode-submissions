/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	val := preorder[0]
	node := &TreeNode{Val: val}

	var mid int
	for i, v := range inorder {
		if val == v {
			mid = i
		}
	}	

	node.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	node.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return node
}
