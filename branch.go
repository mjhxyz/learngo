package main

import (
	"fmt"
	"os"
)

func main() {
	const filename = "abc.txt"
	if contents, err := os.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
