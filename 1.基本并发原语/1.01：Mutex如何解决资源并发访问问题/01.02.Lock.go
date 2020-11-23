package main

import (
	"fmt"
	"sync"
)

/**
使用Mutex情况下10个goroutine进行对同一个变量的计数
*/
func main() {
	for i := 0; i < 10; i++ { //简单for循环
		fmt.Printf("第%d次输出: ", i)
		Lock()
	}
}

//用Mutex处理过的方法
func Lock() {
	// 互斥锁保护计数器
	var mu sync.Mutex
	// 计数器的值
	var count = 0

	// 辅助变量，用来确认所有的goroutine都完成
	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个gourontine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加10万次
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
