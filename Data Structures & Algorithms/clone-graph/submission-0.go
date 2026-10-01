/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	nodes := map[int]*Node{}

	var clone func(node *Node) *Node 
	clone = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		// fmt.Printf("cloning %d\n", node.Val)
		if _, found := nodes[node.Val]; found {
			return nodes[node.Val]
		}

		newNode := &Node{
			Val: node.Val,
			Neighbors: []*Node{},
		}
		nodes[node.Val] = newNode

		for _, n := range node.Neighbors {
			// fmt.Printf("edge %d -> %d\n", node.Val, n.Val)

			cloned := clone(n)
			if cloned != nil {
				newNode.Neighbors = append(newNode.Neighbors, cloned)
			}
		}

		// fmt.Printf("cloned node %d with neighbors %v\n", node.Val, len(newNode.Neighbors))

		return newNode
	}

	return clone(node)
}
