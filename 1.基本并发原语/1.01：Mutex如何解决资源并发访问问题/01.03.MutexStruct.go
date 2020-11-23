package main

import (
	"fmt"
	"sync"
)

/**
有时候，我们还可以采用嵌入字段的方式。通过嵌入字段，你可以在这个 struct 上直接调用 Lock/Unlock 方法。
如果嵌入的 struct 有多个字段，我们一般会把 Mutex 放在要控制的字段上面，然后使用空格把字段分隔开来。
甚至，你还可以把获取锁、释放锁、计数加一的逻辑封装成一个方法，对外不需要暴露锁等逻辑
*/

func main() {
	// 封装好的计数器
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 执行10万次累加
			for j := 0; j < 100000; j++ {
				counter.Incr() // 受到锁保护的方法
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}

// 线程安全的计数器类型
type Counter struct {
	CounterType int
	Name        string

	mu    sync.Mutex
	count uint64
}

// 加1的方法，内部使用互斥锁保护
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 得到计数器的值，也需要锁保护
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
