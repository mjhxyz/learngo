package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6] // [2, 3, 4, 5]
	fmt.Println("arr[2:6] =", s)
	s1 := arr[:6] // [0, 1, 2, 3, 4, 5]
	fmt.Println("arr[:6] =", s1)
	s2 := arr[2:] // [2, 3, 4, 5, 6, 7]
	fmt.Println("arr[2:] =", s2)
	s3 := arr[:] // [0, 1, 2, 3, 4, 5, 6, 7]
	fmt.Println("arr[:] =", s3)

	fmt.Println("After updateSlice(s)")
	updateSlice(s)
	fmt.Println(s)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("Reslice")
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arrx := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1x := arrx[2:6]
	s2x := s1x[3:5]
	fmt.Printf("s1x=%v, len(s1x)=%d, cap(s1x)=%d\n",
		s1x, len(s1x), cap(s1x))
	fmt.Printf("s2x=%v, len(s2x)=%d, cap(s2x)=%d\n",
		s2x, len(s2x), cap(s2x))

	s4 := append(s2, 10)
	s5 := append(s3, 11)
	s6 := append(s4, 12)
	fmt.Println("s4, s5, s6 =", s4, s5, s6)
	fmt.Println("arr =", arr)
}
