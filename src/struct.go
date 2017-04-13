/*
  如果x可获取地址,并且&x的方法中包含了m()
  var x Job
  x.m()
  也是可以的 虽然x不是指针
  这里go 会查找Job类型变量 x 的方法列表,没有就会查找*Job的方法列表 并且将其转化为(&x).m()
  最好使用指针的方法
*/
package main

import "fmt"

type Job struct {
	title string "JD"
	rank  int
}

type JobWithoutTag struct {
	title string "JD"
	rank  int
}

type Employee struct {
	int
	age                 int
	firstName, lastName string
	pack                struct {
		salary int
		stock  int
	}
	*Job
}

func main() {
	job1 := Job{
		"CEO",
		99,
	}

	job2 := Job{
		rank:  100,
		title: "BOSS",
	}

	job3 := Job{
		rank: 200,
	}

	job4 := Job{
		"MANAGER",
		111,
	}

	employee1 := Employee{
		age:       25,
		firstName: "Peng",
		lastName:  "Daimin",
	}

	employee1.pack.salary = 1000
	employee1.pack.stock = 1000

	job5 := Job{
		"STAFF",
		1000,
	}

	// job6 := JobWithoutTag{
	// 	"Bob",
	// 	20,
	// }
	//
	// job7 := JobWithoutTag{}
	//
	// job8 := JobWithoutTag{
	// 	"Bobo",
	// 	20,
	// }

	job9 := Job{}

	job9 = job9

	employee1.Job = &job5

	fmt.Printf("%v\n", job1)

	job1.Help()

	job1.What()

	job2.What()

	job3.What()

	job4.what()

	employee1.Help()

	employee1.Job.Help()

	employee1.What()

	job9.What()
}

func (Job) Help() {
	fmt.Printf("%s\n", "use j.What() to get the j detail")
}

func (job Job) What() {
	fmt.Printf("%v, %v\n", job.title, job.rank)
}

func (job *Job) what() {
	fmt.Printf("%v %v\n", job.rank, job.title)
}

func (Employee) Help() {
	fmt.Printf("%v\n", "use e.what() to get the e detail")
}
