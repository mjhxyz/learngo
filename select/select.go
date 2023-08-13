package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func worker(id int, c chan int) {
	for n := range c { // 等待 c 发送完毕被关闭
		time.Sleep(time.Second)
		fmt.Printf("Worker %d 收到了 %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	c1, c2 := generator(), generator()
	var worker = createWorker(0)

	var values []int                   // 队列存储结果, 防止结果被覆盖掉
	tm := time.After(10 * time.Second) // 10s 后会往这个 channel 送一个时间, 整体的时间
	tick := time.Tick(time.Second)     //定时器，每隔指定时间，都会送值来

	for {
		var activateWorker chan<- int
		var activateValue int
		if len(values) > 0 {
			activateWorker = worker
			activateValue = values[0]
		}

		// 加上 default 类似于非阻塞的逻辑
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activateWorker <- activateValue: // 如果 activateWorker 是 nil, 则不会到这里
			values = values[1:]
		case <-time.After(800 * time.Millisecond): // 每次 select 花的时间
			fmt.Println("获取超时了!!!")
		case <-tick: // 定时器，每隔一段时间就会有值
			// 但是会和上面的 800ms 超时的重叠，导致不太好达到超时的情况
			fmt.Println("队列长度为:", len(values))
		case <-tm: // 这里是从程序运行开始计算的时间
			fmt.Println("时间到了, 结束")
			return
		}
	}
}
