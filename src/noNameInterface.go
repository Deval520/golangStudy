package main

type Students struct {
	name  string
	class int
}

type Monster struct {
	Students
	subject string
}

type MonsterIn interface {
	getSDetails() (string, int)
}

func main() {

}

func (m *Monster) getSDetails() (n string, a int) {
	n = m.Students.name
	a = m.Students.class
	return
}
