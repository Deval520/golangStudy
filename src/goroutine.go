package main

import (
	"fmt"
	"time"
)

func ready(s string, i int) {
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Printf("%v\n", s)
}

func main() {
	go ready("等1秒", 1)
	go ready("等2秒", 2)
	fmt.Printf("%v\n", "等待中")
	//如果不等待goroutine的执行 程序立刻停止 任何正在执行的goroutine都会停止
	time.Sleep(5 * time.Second)
}
