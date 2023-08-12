package tree

import "fmt"

func (node *Node) Traverse() {
	node.TraversaFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraversaFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraversaFunc(f)
	f(node)
	node.Right.TraversaFunc(f)
}
