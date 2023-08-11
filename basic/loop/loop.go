package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	// 对2取模, 余数放在最后, 然后除以2 .... 直到商为0
	result := ""
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	// 一行一行读取文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	// for 条件 === while 条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for { // 死循环
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1011 --> 1101
		convertToBin(72387885),
		convertToBin(0),
	)

	printFile("abc.txt")
	forever()
}
