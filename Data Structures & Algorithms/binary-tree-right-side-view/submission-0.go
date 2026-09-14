/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	current := []*TreeNode{root}
	result := [][]int{}
	for len(current) > 0 {
		level := []int{}
		newCurrent := []*TreeNode{}
		for _, n := range current {
			level = append(level, n.Val)
			if n.Left != nil {
				newCurrent = append(newCurrent, n.Left)
			}
			if n.Right != nil {
				newCurrent = append(newCurrent, n.Right)
			}
		}		
		result = append(result, level)
		current = newCurrent
	}

	return result
}
func rightSideView(root *TreeNode) []int {
	levels := levelOrder(root)
   	result := []int{}

   	for _, level := range levels {
		result = append(result, level[len(level) - 1])
   	}

	return result
}
