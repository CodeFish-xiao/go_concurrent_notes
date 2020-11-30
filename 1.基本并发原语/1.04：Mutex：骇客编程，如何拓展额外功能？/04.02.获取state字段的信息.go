package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const ( //能运行，无视报错即可
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type MyMutex struct {
	sync.Mutex
}

//等待者的数量
func (m *MyMutex) Count() int {
	// 获取state字段的值
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v >> mutexWaiterShift //得到等待者的数值
	v = v + (v & mutexLocked) //再加上锁持有者的数量，0或者1
	return int(v)
}

// 锁是否被持有
func (m *MyMutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}

// 是否有等待者被唤醒
func (m *MyMutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

// 锁是否处于饥饿状态
func (m *MyMutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
}

func count() {
	var mu MyMutex
	for i := 0; i < 1000; i++ { // 启动1000个goroutine
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}

	time.Sleep(time.Second)
	// 输出锁的信息
	fmt.Printf("waitings: %d, isLocked: %t, woken: %t,  starving: %t\n", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}

func main() {
	count()
}
