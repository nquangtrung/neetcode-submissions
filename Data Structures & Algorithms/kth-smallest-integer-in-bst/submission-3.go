/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	if root.Left != nil {
		result = append(result, traverse(root.Left)...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, traverse(root.Right)...)
	}

	return result
}

func dfs(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	curr := root
	count := 0
	for curr != nil || len(stack) > 0 {
		// Push all left node of the current root to stack
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// Handle the left most node on the stack
		n := len(stack) - 1
		curr = stack[n]
		stack = stack[:n]

		// Found the k left most node
		count += 1
		if count == k {
			return curr.Val
		}

		// Handle the right node
		curr = curr.Right
	}

	return -1
}


func kthSmallest(root *TreeNode, k int) int {
	return dfs(root, k)
}