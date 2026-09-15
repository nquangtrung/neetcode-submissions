/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
   return Codec{} 
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	stack := []*TreeNode{root}
	var builder strings.Builder
	for len(stack) > 0 {
		n := len(stack) - 1
		top := stack[n]
		stack = stack[:n]

		if top == nil {
			builder.WriteString("N")
			builder.WriteString(",")
			continue
		}  

		builder.WriteString(strconv.Itoa(top.Val))
		builder.WriteString(",")

		// write left later to be handled first
		stack = append(stack, top.Right, top.Left)	
	}

	return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// fmt.Printf("data: %v \n", data)
	strs := strings.Split(data, ",")
	strs = strs[:len(strs) - 1]

	root := &TreeNode{}
	stack := []*TreeNode{root}
	for i := 0; i < len(strs); {
		// fmt.Printf("new loop %d\n", i)
		for strs[i] != "N" {
			// fmt.Printf("reading %d:%s as Left\n", i, strs[i])
			val, _ := strconv.Atoi(strs[i])
			node := &TreeNode{
				Val: val,
			}
			i += 1

			stack[len(stack) - 1].Left = node
			stack = append(stack, node)
		}

		// Reach the first N as left
		// So this Left is nil, the next strs[i] should be Right
		// fmt.Printf("reading %d:%s as Left\n", i, strs[i])
		i += 1
		for i < len(strs) && strs[i] == "N" {
			// fmt.Printf("reading %d:%s as Right\n", i, strs[i])
			// This is a nil Right so we pop the stack until we meet a right with value
			stack = stack[:len(stack) - 1]
			i += 1
		}

		if i == len(strs) {
			// We read the last N
			break
		}

		// fmt.Printf("reading %d:%s as Right\n", i, strs[i])
		val, _ := strconv.Atoi(strs[i])
		node := &TreeNode{
			Val: val,
		}
		// pop the stack 
		stack[len(stack) - 1].Right = node
		stack = stack[:len(stack) - 1]
		// and then push node into stack
		stack = append(stack, node)
		i += 1
	}

	return root.Left
}
