package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string "nameptr"
	age  int
}

func (p Person) get() {
	fmt.Println("ok")
}

func main() {
	//查看结构体Person的tag
	per := Person{}
	//reflect.TypeOf(i interface{}) 返回这个接口类型对应的reflect.Type对象
	//通过Type提供的方法可以获得这个借口实际的静态类型
	t := reflect.TypeOf(per)
	fmt.Println(t)
	fmt.Println(t.Field(0).Tag)
	fmt.Println(t.Name())
	fmt.Println(t.Align())
	fmt.Println(t.NumMethod())
	fmt.Println(t.String())
	fmt.Println(t.NumField())
}
