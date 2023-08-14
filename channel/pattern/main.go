package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成消息的生成器
func msgGen(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(5000)) * time.Millisecond):
				c <- fmt.Sprintf("服务 %s: 消息 %d", name, i)
			case <-done:
				fmt.Println("正在做清理工作")
				time.Sleep(2 * time.Second) // 2s 清理
				fmt.Println("清理工作完成")
				done <- struct{}{} // 双向的 done, 表示清理完成, 能退出
				return
			}
			i++
		}
	}()
	return c
}

func fanIn(chs ...chan string) chan string {
	c := make(chan string)
	// 创建多个 goroutine, 分别从不同的 channel 中获取数据
	for _, ch := range chs {
		go func(ch chan string) {
			for {
				c <- <-ch
			}
		}(ch)
	}
	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- <-c1:
			case c <- <-c2:
			}
		}
	}()
	return c
}

// 非阻塞等待,(消息, 是否有数据)
func nonBlockingWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

// 超时等待
func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		return "", false
	}
}

func main() {
	done := make(chan struct{})
	m1 := msgGen("service1", done)
	for i := 0; i < 5; i++ {
		if m, ok := timeoutWait(m1, 2*time.Second); ok {
			fmt.Println(m)
		} else {
			fmt.Println("m1 超时")
		}
	}
	done <- struct{}{}
	<-done
}
