package main

import (
	"fmt"
	"sync"
)

func foo1() {
	var mu sync.Mutex
	defer mu.Unlock() //注意！之前没有lock哦
	fmt.Println("hello world!")
}
func main() {
	foo1()
}
