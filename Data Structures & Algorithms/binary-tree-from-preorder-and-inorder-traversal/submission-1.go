/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func buildTree(preorder []int, inorder []int) *TreeNode {
    // Array lookup instead of map allocation (handles values -3000 to 3000)
    var idxMap [6001]int
    for i, val := range inorder {
        idxMap[val+3000] = i
    }

    preIdx := 0
    var helper func(left, right int) *TreeNode
    helper = func(left, right int) *TreeNode {
        if left > right {
            return nil
        }

        rootVal := preorder[preIdx]
        preIdx++

        root := &TreeNode{Val: rootVal}
        mid := idxMap[rootVal+3000]

        root.Left = helper(left, mid-1)
        root.Right = helper(mid+1, right)
        return root
    }

    return helper(0, len(inorder)-1)
} 
