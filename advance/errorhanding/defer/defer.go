package main

import (
	"bufio"
	"fmt"
	fib "learngo/advance/functional"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("this is panic")
	fmt.Println(4)
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	//err = errors.New("这是一个自定义错误")
	if err != nil {
		// OpenFile: // If there is an error, it will be of type *PathError.
		if pathError, ok := err.(*os.PathError); !ok {
			// 不是 PathError:
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err,
			)
		}
	}
	defer file.Close()

	// 用缓冲区写入文件
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
