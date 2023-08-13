package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {
	// 判断 channel 被关闭方法2
	for n := range w.in { // 等待 c 发送完毕被关闭
		fmt.Printf("Worker %d 收到了 %c\n", id, n)
		w.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

type worker struct {
	in   chan int
	done func()
}

func channelDemo() {
	var workers [10]worker
	var wg sync.WaitGroup

	for i, _ := range workers {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20) // 添加 20 个任务
	for i, worker := range workers {
		worker.in <- 'a' + i // 向 channel 发送任务数据
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait() // 等待 20 个任务做完
}

func main() {
	channelDemo()
}
