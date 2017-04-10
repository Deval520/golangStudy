package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	i    int
	data [10]int
}

func main() {
	// slice1 := make([]float64, 5)
	// slice1 = []float64{10.2, 443.6532, 434.8978, 56.4, 787.6}
	// fmt.Println(avg(slice1))
	//
	// fmt.Println(sort(2, 4, 1, 46))
	//
	// var s stack
	// fmt.Println(s.i)
	// s.push(10)
	// fmt.Println(s, s.i)
	//
	// s.push(20)
	// fmt.Println(s.pop())
	// fmt.Println(s.i)
	// fmt.Println(s.pop())
	// fmt.Println(s.i)
	// fmt.Println(s.String())

	// // fmt.Println(getFeibo(5))
	// test := func(i int) int {
	// 	return i * i
	// }
	crr := []int{2, 4, 4, 5}
	// fmt.Println(map(test,crr))
	//
	// test1 := func(i int) int {
	// 	return i * i
	// }
	fmt.Println(Maptest(test, crr))

}

func avg(arr []float64) (ave float64) {
	var sum float64
	for _, val := range arr {
		sum += val
	}
	ave = sum / float64(len(arr))
	return
}

func sort(arg ...int) []int {
	for i := 0; i < len(arg); i++ {
		for j := i + 1; j < len(arg); j++ {
			if arg[i] < arg[j] {
				tem := arg[i]
				arg[i] = arg[j]
				arg[i] = tem
			}
		}
	}
	return arg
}

func (s *stack) push(v int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = v
	s.i++
	// fmt.Println(s.data)
}

func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

func (s stack) String() (str string) {
	for i := 0; i < len(s.data); i++ {
		str += "[" + strconv.Itoa(i) + " : " + strconv.Itoa(s.data[i]) + "]"
	}
	return
}

func Map(f func(int) int, crr []int) []int {
	//f 返回原数组元素的平方
	brr := make([]int, len(crr))

	for k, v := range crr {
		brr[k] = f(v)
	}

	return brr
}

func test(i int) (c int) {
	c = i * i
	return
}

func Maptest(f func(int) int, arr []int) []int {
	j := make([]int, len(arr))
	for k, v := range arr {
		j[k] = f(v)
	}
	return j
}
