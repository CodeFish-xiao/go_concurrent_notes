package main

import (
	"fmt"
	"strings"
	"sync"
)

type Queue struct {
	cond *sync.Cond
	data []interface{}
	capc int
	logs []string
}

func NewQueue(capacity int) *Queue {
	return &Queue{cond: &sync.Cond{L: &sync.Mutex{}}, data: make([]interface{}, 0), capc: capacity, logs: make([]string, 0)}
}

func (q *Queue) Enqueue(d interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for len(q.data) == q.capc {
		q.cond.Wait()
	}
	// FIFO入队
	q.data = append(q.data, d)
	// 记录操作日志
	q.logs = append(q.logs, fmt.Sprintf("En %v\n", d))
	// 通知其他waiter进行Dequeue或Enqueue操作
	q.cond.Broadcast()

}

func (q *Queue) Dequeue() (d interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for len(q.data) == 0 {
		q.cond.Wait()
	}
	// FIFO出队
	d = q.data[0]
	q.data = q.data[1:]
	// 记录操作日志
	q.logs = append(q.logs, fmt.Sprintf("De %v\n", d))
	// 通知其他waiter进行Dequeue或Enqueue操作
	q.cond.Broadcast()
	return
}

func (q *Queue) Len() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return len(q.data)
}

func (q *Queue) String() string {
	var b strings.Builder
	for _, log := range q.logs {
		//fmt.Fprint(&b, log)
		b.WriteString(log)
	}
	return b.String()
}
