package main

import (
	"fmt"
	//可以使用包的别名导入
	two "one"
)

func main() {
	a := 1
	fmt.Printf("%v, %v, %T\n", two.Even(a), a, a)
	duoCan("abc", 1, 3, 5, 5)

}

//变参函数
func duoCan(a string, b ...int) {
	fmt.Printf("%v ,%v\n", a, b)
}
