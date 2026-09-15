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

func dfs(root * TreeNode) []int {
	stack := []*TreeNode{root}
	visited := map[*TreeNode]bool{}
	result := []int{}
	for len(stack) > 0 {
		top := stack[len(stack) - 1]
		if top.Left != nil && !visited[top.Left] {
			stack = append(stack, top.Left)
			continue
		}

		result = append(result, top.Val)
		stack = stack[:len(stack) - 1]
		visited[top] = true

		if top.Right != nil && !visited[top.Right] {
			stack = append(stack, top.Right)
		}
	}

	return result
}


func kthSmallest(root *TreeNode, k int) int {
	arr := dfs(root)

	return arr[k - 1]
}