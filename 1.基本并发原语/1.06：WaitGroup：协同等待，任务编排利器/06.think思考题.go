package main

import (
	"sync"
	"unsafe"
)

/*
注意点：
1. 这里用了一个函数来实现，更常见的可以自己封一个类。用函数实现时注意用指针传递wg
2. 返回的两个值分别是state和wait，state是要完成的waiter计数值（即等待多少个goroutine完成）；wait是指有多少个sync.Wait在等待（和前面的waiter不是一个概念）。
*/
func getStateAndWait(wgp *sync.WaitGroup) (uint32, uint32) {
	var statep *uint64
	if uintptr(unsafe.Pointer(wgp))%8 == 0 {
		statep = (*uint64)(unsafe.Pointer(wgp))
	} else {
		statep = (*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(wgp)) + unsafe.Sizeof(uint32(0))))
	}
	return uint32(*statep >> 32), uint32(*statep)
}
