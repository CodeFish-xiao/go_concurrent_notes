package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	wg.Add(-10) //将-10作为参数调用Add，计数值被设置为0

	wg.Add(-1) //将-1作为参数调用Add，如果加上-1计数值就会变为负数。这是不对的，所以会触发panic
}
