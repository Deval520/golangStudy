package main

import "fmt"

func main() {
	//对一个数组进行排序 并返回数组的所有奇数
	var arr [5]int = [5]int{567, 5, 89, 98, 4}
	//数组指针
	var ptr *[5]int = &arr

	brr := funcOne(ptr)
	fmt.Println(arr, brr)

	funcTwo()

	funcThree(1, 2, 34, 5)

	//函数也是值 func()类型的值
	a := func() {
		fmt.Println("a")
	}

	func() {
		fmt.Println("over")
	}()

	a()
	//打印a类型为func()
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", arr)
	fmt.Printf("%T\n", funcSeven)

	c := 10
	callback(c, funcSix)

}

func funcOne(ptr *[5]int) (arr []int) {
	//返回一个存放奇数的切片
	for i := 0; i < len(ptr); i++ {
		if (*ptr)[i]%2 != 0 {
			arr = append(arr, (*ptr)[i])
		}
		for j := i + 1; j < len(ptr); j++ {
			if (*ptr)[i] < (*ptr)[j] {
				tem := (*ptr)[i]
				(*ptr)[i] = (*ptr)[j]
				(*ptr)[j] = tem
			}
		}
	}
	//函数即将推出前执行 可以访问任何命名返回参数 闭包?
	defer func(x int) {
		fmt.Println(x)
	}(10)

	return
}

func funcTwo() {
	//延迟的函数式按照后进先出的顺序执行的
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func funcThree(arg ...int) {
	//参数arg是一个int切片 变参函数
	fmt.Println(arg)
}

func funcSix(b int) {
	// a := "funcsix"
	fmt.Println("callback")
	// return a

}

func funcSeven(b int) (a string) {
	a = "funcseven"
	return
}

//函数作为参数传递 需要?
func callback(y int, f func(int)) {
	// a := f(y)
	f(y)
}
