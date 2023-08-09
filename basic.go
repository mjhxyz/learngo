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
	var c int
	c = int(math.Sqrt(float64(a*b + b*b)))
	fmt.Println(c)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()
}
