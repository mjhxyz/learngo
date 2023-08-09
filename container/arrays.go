package main

import "fmt"

func printArray(arr [5]int) {
	// 值传递
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}          // := 需要给初始值
	arr3 := [...]int{2, 4, 6, 8, 10} // ... 表示根据初始值自动推断长度
	var grid [4][5]int               // 二维数组
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	printArray(arr1)
	//printArray(arr2)
	printArray(arr3)

	fmt.Println(arr1)
	fmt.Println(arr3)
}
