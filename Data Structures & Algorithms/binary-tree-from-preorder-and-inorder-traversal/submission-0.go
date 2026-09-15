/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func buildTree(preorder []int, inorder []int) *TreeNode {
	// Map to store value -> index for O(1) lookups in inorder array
	inorderMap := make(map[int]int)
	for i, val := range inorder {
		inorderMap[val] = i
	}

	preIdx := 0

	var helper func(left, right int) *TreeNode
	helper = func(left, right int) *TreeNode {
		// Base case: no elements left to construct subtree
		if left > right {
			return nil
		}

		// Select current root from preorder traversal
		rootVal := preorder[preIdx]
		preIdx++

		root := &TreeNode{Val: rootVal}

		// Root index splits inorder array into left and right subtrees
		mid := inorderMap[rootVal]

		// Build left subtree first because preorder is (Root -> Left -> Right)
		root.Left = helper(left, mid-1)
		root.Right = helper(mid+1, right)

		return root
	}

	return helper(0, len(inorder)-1)
} 
