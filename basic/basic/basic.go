package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 包内部变量，不是全局变量!!!!!
var aa = 3
var ss = "kkk"
var bb = true

// 还可放在一个括号里面
var (
	cc = 3
	dd = "kkk"
	ee = true
)

// 函数外面不能使用 := 定义变量
// bb := true

func variableZeroValue() {
	// 变量名在类型之前
	var a int
	var s string
	// %q 打印字符串,能把双引号打印出来
	fmt.Printf("%d %q\n", a, s)
}

func variableInitValue() {
	// 可以一行定义多个变量
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	// 可以直接推断类型
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	// 省略 var, 第一次使用变量的时候必须使用 :=
	// var 能不用就不用
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func euler() {
	// 复数表示
	fmt.Printf("%.3f\n",
		cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	const (
		filename2 = "dcd.txt"
		e, f      = 5, 6
	)
	var c int
	// 常量没有说明类型的时候，可以当作字面量
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c, filename2, e, f)
}

func enums() {
	// 没有特殊的关键字，用一组 const 定义表示
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	// 还可以简写
	const (
		cpp2 = iota
		_    // 跳过一个
		java2
		python2
		golang2
	)
	// 还有 b, kb, mb, gb, tb, pb
	// iota 能参与运算
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(cpp2, java2, python2, golang2)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()
	consts()
	enums()
}
