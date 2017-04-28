package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
		timeout <- false
		// timeout <- false
	}()
	ch := make(chan int, 0)
	select {
	case <-ch:
	case <-timeout:
		fmt.Printf("%v\n", "timeout!")
	}
	//select会一直等到某-个case语句完成，则select语句结束

	//利用 default 检测channel是否满
	ch2 := make(chan int, 1)
	// ch2 <- 100
	// ch2 <- 300
	select {
	case ch2 <- 200:
		fmt.Printf("%v\n", "插入200")
	default:
		fmt.Printf("%v\n", "channel已满")
	}
}
