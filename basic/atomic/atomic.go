package main

import (
	"fmt"
	"sync"
	"time"
)

type AtomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *AtomicInt) inc() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *AtomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Lock()
	return a.value
}

func main() {
	var a AtomicInt
	a.inc()
	go func() {
		a.inc()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
