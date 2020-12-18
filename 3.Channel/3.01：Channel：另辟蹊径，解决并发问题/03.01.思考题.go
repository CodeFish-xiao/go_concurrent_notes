package main

import (
	"fmt"
	"time"
)

type NumberChan struct {
	Ch            chan int
	ChannelNumber int
}

func (nch *NumberChan) SendNotify() {
	go func() {
		nch.Ch <- nch.ChannelNumber
	}()
}

func (nch *NumberChan) PrintInfo() {
	fmt.Println(nch.ChannelNumber)
	time.Sleep(time.Second)
}

func NewNumberChan(seq int) *NumberChan {
	nch := NumberChan{
		Ch:            make(chan int),
		ChannelNumber: seq,
	}
	return &nch
}

func main() {
	var (
		nch1 = NewNumberChan(1)
		nch2 = NewNumberChan(2)
		nch3 = NewNumberChan(3)
		nch4 = NewNumberChan(4)
	)
	go func() {
		nch1.SendNotify()
	}()
	for {
		select {
		case <-nch1.Ch:
			nch1.PrintInfo()
			nch2.SendNotify()
		case <-nch2.Ch:
			nch2.PrintInfo()
			nch3.SendNotify()
		case <-nch3.Ch:
			nch3.PrintInfo()
			nch4.SendNotify()
		case <-nch4.Ch:
			nch4.PrintInfo()
			nch1.SendNotify()
		}
	}

}
