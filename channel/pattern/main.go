package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成消息的生成器
func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("服务 %s: 消息 %d", name, i)
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

func main() {
	m1 := msgGen("service1")
	m2 := msgGen("service2")
	m3 := msgGen("service3")
	m := fanIn(m1, m2, m3)
	//m := fanInBySelect(m1, m2)
	for {
		fmt.Println(<-m)
	}
}
