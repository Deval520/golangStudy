package main

import "fmt"

type MenInter interface {
	GetDetails() (n string)
}

type Men struct {
	int
	name string
	age  int
}

func main() {
	//得到缺省值int
	men := Men{
		// num:  100,
		age:  25,
		name: "deval",
	}

	slice1 := make([]MenInter, 1)
	slice1[0] = &men
	fmt.Printf("%v\n", slice1[0].GetDetails())
	for _, v := range slice1 {
		fmt.Printf("%v\n", v.GetDetails())
	}
}

func (m *Men) GetDetails() string {
	return m.name
}
