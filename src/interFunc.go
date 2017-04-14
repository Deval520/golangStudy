// package main
//
// import "fmt"
//
// type Empty interface{}
//
// type Stack struct {
// 	data  [10]int
// 	index int
// }
//
// type S interface {
// 	Push(int)
// 	Pop() int
// }
//
// type stack struct {
// 	data  [10]int
// 	index int
// }
//
// type SS interface {
// 	push(int)
// 	pop() int
// }
//
// func main() {
// 	var s S
// 	ss := make([]S, 1)
// 	stack := Stack{}
// 	stack.Push(10)
// 	// ss[0] = &stack
// 	s = &stack
// 	get(s)
// 	// fmt.Printf("%v\n", ss)
// }
//
// //指定了接口的类型是指针类型
// func (s *Stack) Push(i int) {
// 	s.data[s.index] = i
// 	s.index++
// }
//
// func (s *Stack) Pop() int {
// 	s.index--
// 	return s.data[s.index]
// }
//
// func (s stack) push(i int) {
// 	s.data[s.index] = i
// 	s.index++
// }
//
// func (s stack) pop() (i int) {
// 	s.index--
// 	i = s.data[s.index]
// 	return
// }
//
// func get(s Empty) {
// 	switch s.(type) {
// 	case *Stack:
// 		fmt.Printf("%v\n", s.Pop())
// 	case stack:
// 		fmt.Printf("%v\n", s.pop())
// 	}
// }
//
package main

import "fmt"

type Stack1 struct {
	index int
	data  [5]int
}

type Stack2 struct {
	index int
	data  [5]int
}

type sinter interface {
	Get() int
	Push(int)
}

type empty interface{}

//
func (s *Stack1) Get() int {
	if s.index < 0 {
		return 0
	}
	s.index--
	return s.data[s.index]
}

func (s *Stack1) Push(k int) {
	if s.index > 9 {
		return
	}

	s.data[s.index] = k
	s.index++
}

//
func (s2 Stack2) Get() int {
	if s2.index < 0 {
		return 0
	}
	s2.index--
	return s2.data[s2.index]
}

func (s2 Stack2) Push(k int) {
	if s2.index > 9 {
		return
	}
	s2.data[s2.index] = k
	s2.index++
}

func f(s sinter) {
	switch s.(type) {
	case *Stack1:
		fmt.Printf("%v\n", "*Stack1 type")
	case Stack2:
		fmt.Printf("%v\n", "Stack2 type")

	}
}

func main() {
	// var ss sinter
	// var s interface{}
	var sss interface{}
	s := make([]sinter, 2)
	s1 := Stack1{}
	s[0] = &s1
	s2 := Stack2{}
	s2.Push(3)
	s1.Push(2)
	s[1] = s2
	f(s[0])
	f(s[1])
	sss = &s1

	// fmt.Printf("%v\n", s[0].Get())
	value := sss.(sinter)

	fmt.Printf("%v\n", value)
	// value := s[0].(sinter)
	// fmt.Printf("%v \n", value.Get())
}
