package tree

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	// 左中右
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
