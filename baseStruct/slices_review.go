package main

import "fmt"

func main() {
	//数组定长 切片长度可变、容量固定相同元素序列
	//长度可变? 容量固定?
	slice1 := make([]int, 5, 10) //申明一个切片
	arr := make([]int, 5)
	brr := []int{}

	//切片本质引用同一个数组
	slice2 := []string{"a", "b", "c", "d", "e"}
	slice3 := slice2[2 : len(slice2)-1]
	slice4 := slice2[:3]
	slice3[0] = "8"

	arr = append(arr, 1, 3, 4, 5, 6, 5)

	slice1 = append(slice1, 1, 2, 3)

	fmt.Println(slice1)

	fmt.Println(len(slice1))

	fmt.Println(cap(slice1))

	fmt.Println(arr)

	fmt.Println(len(arr))

	fmt.Println(cap(arr))

	fmt.Println(brr)

	fmt.Println(len(brr))

	fmt.Println(cap(brr))

	fmt.Println(slice2, slice3, slice4)
}
