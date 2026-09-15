/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// Linear scan using pointer math / index search
	// On tiny slices (N <= 3000), CPU cache vectorization makes this instant
	mid := 0
	for i, v := range inorder {
		if v == rootVal {
			mid = i
			break // Critical fix: break early!
		}
	}

	// Slice indexing passes existing memory views without heap allocation
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}