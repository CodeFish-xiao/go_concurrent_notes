package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func GoID() int {
	var buf [64]byte
	//通过 runtime.Stack 方法获取栈帧信息
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
func main() {

	go func() { fmt.Println("我是大哥", GoID()) }()
	go func() {
		fmt.Println("我是二哥", GoID())
	}()
	go func() {
		fmt.Println("我是三弟", GoID())
	}()

	fmt.Println("今天我们桃园三结义", GoID())
	time.Sleep(5000)
}
