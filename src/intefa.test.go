package main

import "fmt"

type Men struct {
	name string
	age  int
}

//abstract class
type People interface {
	GetDetails() (name string, age int)
}

func main() {
	men := Men{
		"deval",
		25,
	}

	var a interface{}
	i := 1
	a = i
	people := make([]People, 1)
	people[0] = &men

	name, _ := men.GetDetails()
	fmt.Printf("%v\n", name)
	fmt.Printf("%T\n", people)
	fmt.Printf("%v\v", people)
	n, _ := people[0].GetDetails()
	fmt.Printf("%v\n", n)
	fmt.Printf("%v\n", a)
	fmt.Println(a)

}

func (m *Men) GetDetails() (name string, age int) {
	name = m.name
	age = m.age
	return
}
