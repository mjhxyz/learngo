package main

import (
	"learngo/tree"
)

type MyTreeNode struct {
	*tree.Node // Embeding
}

func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	// 左右中
	node := MyTreeNode{myNode.Left}
	node.postOrder()
	right := MyTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}

func main() {

	// 创建结构体的几种方式
	root := MyTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	root.postOrder()
}
