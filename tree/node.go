package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) SetValue(value int) {
	// 指针接收者
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.Value = value
}

func (node Node) Print() {
	// 给结构定义方法
	fmt.Print(node.Value, " ")
}

func CreateNode(value int) *Node {
	// 返回局部变量的地址!!!
	// 不需要知道
	return &Node{Value: value}
}
