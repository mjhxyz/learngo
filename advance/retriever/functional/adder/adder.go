package main

import "fmt"

func adder() func(int) int {
	s := 0
	return func(value int) int {
		s += value
		return s
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}
}
