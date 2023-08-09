package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("value=%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating slice...")
	var s []int // Zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, i)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16) // 创建一个长度为16的slice
	printSlice(s2)
	s3 := make([]int, 10, 32) // 创建一个长度为10, 容量为32的slice
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1) // 将s1的内容复制到s2
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...) // 删除s2[3]

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(front, tail)
	printSlice(s2)
}
