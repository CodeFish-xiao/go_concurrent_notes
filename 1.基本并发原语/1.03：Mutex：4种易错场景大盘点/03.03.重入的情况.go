package main

import (
	"fmt"
	"sync"
)

func foo3(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l) //反正就是一个锁不能锁两次啦
	l.Unlock()
}

func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func main() {
	l := &sync.Mutex{}
	foo3(l)
}
