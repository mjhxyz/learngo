package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	// 左中右
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func (node *treeNode) setValue(value int) {
	// 指针接收者
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.value = value
}

func (node treeNode) print() {
	// 给结构定义方法
	fmt.Print(node.value, " ")
}

func createNode(value int) *treeNode {
	// 返回局部变量的地址!!!
	// 不需要知道
	return &treeNode{value: value}
}

func main() {

	// 创建结构体的几种方式
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	root.right.left.setValue(4)

	root.traverse()

	// 使用切片创建
	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
}
