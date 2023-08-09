package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes你好世界啊!"
	fmt.Println(s)
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println(
		"Rune count:",
		utf8.RuneCount([]byte(s)))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) // 解码第一个rune, 返回rune和长度
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 直接使用rune 就完事了
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
}
