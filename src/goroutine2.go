package main

import (
	"fmt"
	"time"
)

//定义一个int 类型的channel(传输整数).这个channel传输整数 全局变量 goroutine可以访问到它
var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Printf("%v is ready\n", w)
	//发送整数1到channel c
	c <- sec
}

func main() {
	//初始化赋值
	// d := make([]int)
	// d = d.append(1)
	c = make(chan int)
	go ready("one", 1)
	go ready("two", 2)
	fmt.Printf("%v\n", "i'm waiting")
	//等待直到从channel c上接受到值
	f := <-c
	l := <-c
	//两个goroutine 收到两个值
	fmt.Printf("%v %v\n", f, l)
	// fmt.Printf("%v\n", d)

}
