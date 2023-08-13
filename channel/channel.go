package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	// 判断 channel 被关闭方法2
	for n := range c { // 等待 c 发送完毕被关闭
		fmt.Printf("Worker %d 收到了 %c\n", id, n)
	}
	// 判断 channel 被关闭方法1
	//for {
	//	//n := <-c // 从 channel 拿出数据
	//	n, ok := <-c // ok: 是否还有值
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d 收到了 %c\n", id, n)
	//}
}

func bufferedChannel() {
	// 设置 channel 的缓冲区大小为 3
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func channelDemo() {
	// chan of int => 是一个 channel 里面的内容是 int
	//var c chan int // c == nil 不能直接用

	//c := make(chan int)       // 这里的 channel 才可以直接用了
	var channels [10]chan<- int // 10 个 channel

	for i := 0; i < 10; i++ { // 启动 10 个 Wroker
		// 启动 10 个 worker 并且分发所有的 channel
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		// 如果发送了一个数据，但是没有其他的协程接收,会死锁的
		channels[i] <- 'a' + i // 向 channel 发送数据
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}
func channelClose() {
	// 设置 channel 的缓冲区大小为 3
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//channelDemo()
	// bufferedChannel()
	channelClose()
}
