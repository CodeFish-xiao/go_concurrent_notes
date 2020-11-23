package main

import (
	"fmt"
	"sync"
)

/**
在没有加锁的情况下10个goroutine进行对同一个变量的计数
*/
func main() {
	for i := 0; i < 10; i++ { //简单for循环
		fmt.Printf("第%d次输出: ", i)
		UnLock()
	}

}

func UnLock() {
	var count = 0
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}
