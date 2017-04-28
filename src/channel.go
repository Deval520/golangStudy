package main

import (
	"fmt"
	"time"
)

func main() {
	//并行执行打印0 - 10

	//无缓冲channel 发送数据到channel会被阻塞 直到有数据会被读取
	ch := make(chan int, 0)
	go prtNum(ch)
	for i := 0; i < 10; i++ {
		//ch 接收数据

		ch <- i
		fmt.Printf("%v\n", "执行main")
	}

}

func prtNum(ch chan int) {
	// for {
	// j := <-ch
	// fmt.Printf("%v\n", j)
	// }
	for {
		j := <-ch
		fmt.Printf("%v\n", j)
		time.Sleep(time.Duration(2) * time.Second)
	}
}
