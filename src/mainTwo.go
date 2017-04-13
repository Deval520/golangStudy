package main

import "fmt"

type nameAge struct {
	name string
	age  int
}

func main() {
	//new make
	/*
	   make 只能创建slice map channel
	*/

	//new 返回的是指针类型
	var p *[]int = new([]int)

	fmt.Printf("%v\n", p == nil)
	//false

	//返回的值类型
	// var v []int = make([]int, 10)

	*p = make([]int, 100)

	fmt.Printf("%v\n", len(*p))

	var deval *nameAge = new(nameAge)
	deval.age = 20
	deval.name = "deval"

	fmt.Printf("%v\n", *deval)
}

//os包 返回一个File类型的指针
// func NewFile(fd int, name string) *File {
// 	if fd < 0 {
// 		return nil
// 	}
// 	return &File{fd, name, 0}
//  复合申明获取地址意味着告诉编译器在堆中分配空间 而不是在栈中
// }
//
