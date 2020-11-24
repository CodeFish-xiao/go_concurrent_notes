package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	Count int
}

func main() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo2(c) // 复制锁
}

// 这里Counter的参数是通过复制的方式传入的
func foo2(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}
