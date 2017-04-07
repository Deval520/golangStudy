package main

import "fmt"

var a int = 5
var b int

func main() { //执行顺序从上至下
	// rec(0)
	b = 6
	// func1()
	fun2()
	fmt.Println(a)
}

//递归函数
func rec(i int) {
	if i == 10 {
		return
	}

	rec(i + 1)
	fmt.Printf("%d ", i)
}

func func1() {
	a := 6
	fmt.Println(a)
}

func fun2() {
	b = 8
	fmt.Println(b)
}
