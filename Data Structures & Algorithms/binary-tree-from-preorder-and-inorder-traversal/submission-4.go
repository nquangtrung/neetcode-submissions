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

	preIdx := 0

	var helper func(left, right int) *TreeNode 
	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		root := &TreeNode{
			Val: preorder[preIdx],
		}
		preIdx += 1

		mid := inorderMap[root.Val]

		root.Left = helper(left, mid - 1)
		root.Right = helper(mid + 1, right)

		return root
	}

	return helper(0, len(inorder) - 1)
}