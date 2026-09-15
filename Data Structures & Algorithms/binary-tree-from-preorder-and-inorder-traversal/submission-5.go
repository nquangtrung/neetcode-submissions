/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := map[int]int{}
	for i, n := range inorder {
		inorderMap[n] = i
	}

	var helper func(preIdx, left, right int) *TreeNode 
	helper = func(preIdx, left, right int) *TreeNode {
		if left > right {
			return nil
		}

		root := &TreeNode{
			Val: preorder[preIdx],
		}

		mid := inorderMap[root.Val]

		root.Left = helper(preIdx + 1, left, mid - 1)
		root.Right = helper(preIdx + mid - left + 1, mid + 1, right)

		return root
	}

	return helper(0, 0, len(inorder) - 1)
}