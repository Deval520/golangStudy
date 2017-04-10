package main

import "fmt"

func main() {
	a := pluxX()
	// fmt.Println(a(8)))
	fmt.Printf("%T\n", a)
	// fmt.Println(a(4)))
	b := a(8)
	fmt.Println(b)
}

//返回一个函数
func pluxX() func(int) int {
	return func(x int) int {
		return x + 2
	}
}
