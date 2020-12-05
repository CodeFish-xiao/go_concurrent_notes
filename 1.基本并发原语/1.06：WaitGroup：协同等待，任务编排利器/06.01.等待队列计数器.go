package main

import (
	"fmt"
	"sync"
	"time"
)

// 线程安全的计数器
type Counter struct {
	mu    sync.Mutex
	count uint64
}

// 对计数值加一
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 获取当前的计数值
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// sleep 1秒，然后计数值加1
func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}

func main() {
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(10) // WaitGroup的值设置为10

	for i := 0; i < 10; i++ { // 启动10个goroutine执行加1任务
		go worker(&counter, &wg)
	}
	// 检查点，等待goroutine都完成任务
	wg.Wait()
	// 输出当前计数器的值
	fmt.Println(counter.Count())
}
