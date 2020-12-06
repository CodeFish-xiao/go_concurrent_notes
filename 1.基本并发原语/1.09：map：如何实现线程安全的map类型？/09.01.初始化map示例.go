package main

import (
	"fmt"
	"time"
)

//未初始化
func main1() {
	var m map[int]int
	m[100] = 100
}

//取0值
func main2() {
	var m map[int]int
	fmt.Println(m[100])
}

//结构体包含忘记初始化
type Counter struct {
	Website      string
	Start        time.Time
	PageCounters map[string]int
}

//结构体包含忘记初始化
func main3() {
	var c Counter
	c.Website = "baidu.com"

	c.PageCounters["/"]++
}

func main() {
	//main1()
	main2()
	main3()
}
