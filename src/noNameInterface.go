package main

import "fmt"

/*
	struct 集成 重载

	interface 存储
*/

type Students struct {
	int
	name  string
	age   int
	class string
}

type StudentsInter interface {
	GetDetails() (n string, a int, c string)
}

type Monster struct {
	//匿名字段
	Students
	subject string
}

func main() {
	slice1 := make([]StudentsInter, 0)
	students1 := Students{
		1,
		"deval",
		25,
		"one",
	}
	monster1 := Monster{
		Students{
			2,
			"val",
			25,
			"two",
		},
		"english",
	}

	slice1 = append(slice1, &students1, &monster1)

	for _, v := range slice1 {
		name, _, _ := v.GetDetails()
		fmt.Printf("%v\n", name)
	}

	fmt.Printf("%v\n", monster1.GetSubject())

}

func (s *Students) GetDetails() (name string, age int, class string) {
	name = s.name
	age = s.age
	class = s.class
	return
}

func (m *Monster) GetSubject() (s string) {
	s = m.subject
	return
}
