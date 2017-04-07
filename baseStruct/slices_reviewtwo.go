package main

import "fmt"

func main() {
	slice0 := make([]string, 4)

	slice0 = []string{"a", "b", "c", "d"}

	slice0 = append(slice0, "e", "f")

	slice1 := slice0[0:5]

	slice1[0] = "b"

	arr0 := [4]int{1, 2, 3, 4}

	arr1 := arr0[:3] //本质也是引用了同一个数组

	arr1[0] = 2

	// brr := [4]int{1, 2, 3, 4}
	//
	// slice2 := make(brr, 4)

	fmt.Println(slice0)

	fmt.Println(len(slice0))

	fmt.Println(cap(slice0))

	fmt.Println(slice1)

	fmt.Println(arr1)

	fmt.Println(arr0)

	// fmt.Println(slice2, len(slice2), cap(slice2))
}
