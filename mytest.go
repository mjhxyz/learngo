package main

import "fmt"

func main() {
	name := "走,忽略"
	for pos, ch := range name {
		fmt.Printf("%d %c\n", pos, ch)
	}
	runes := []rune(name)
	for pos, ch := range runes {
		fmt.Printf("%d %c\n", pos, ch)
	}
}
