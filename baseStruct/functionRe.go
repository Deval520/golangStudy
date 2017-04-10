package main

import "fmt"

func main() {
	// arr := [5]int{4, 31, 89, 4, 0}

	var arr = [5]int{4, 31, 89, 9, 0}
	//指针数组
	var ptr [5]*int
	//数组指针
	var ptr2 *[5]int = &arr
	for k, v := range arr {
		ptr[k] = &v
	}

	for _, v := range ptr {
		fmt.Println(v)
	}
	// fmt.Println(arr)

	// max, crr := calArrMax(*arr)
	// fmt.Println(max, crr)
	calArrMax(ptr2)
	fmt.Println(arr)
	fmt.Println(calArrMax(ptr2))
}

func calArrMax(arr *[5]int) (rr *[5]int) {
	//冒泡排序
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if (*arr)[i] < (*arr)[j] {
				tem := (*arr)[i]
				(*arr)[i] = (*arr)[j]
				(*arr)[j] = tem
			}
		}
	}
	// max = (*arr)[0]
	// crr = arr
	// return max, crr
	// fmt.Println(*arr)
	rr = arr
	//defer 语句在函数即将退出时执行
	defer fmt.Println("over")
	return

}
