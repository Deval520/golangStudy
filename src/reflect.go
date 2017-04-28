package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	name string "name tag"
	age  int
}

func (f Foo) GetAge() (age int) {
	return f.age
}

func main() {
	// i := 10
	// j := "hello"
	// foo := Foo{
	// 	"deval",
	// 	25,
	// }
	//
	// fmt.Printf("%v\n", reflect.TypeOf(i))
	// fmt.Printf("%v\n", reflect.TypeOf(j))
	// fmt.Printf("%v\n", reflect.TypeOf(foo).Field(0))
	//
	// fmt.Printf("%v\n", reflect.TypeOf(foo).NumMethod())
	/*
	  TypeOf:   ValueOf:
	*/
	//类型为struct的初始化
	a := Foo{
		"deval",
		25,
	}

	fmt.Printf("%v\n", reflect.TypeOf(a))

	//类型为指针的初始化
	b := new(Foo)
	b.name = "macholy"
	fmt.Printf("%v\n", *b)
	fmt.Printf("%v\n", reflect.TypeOf(b))

	//----------------------
	c := Foo{
		"macholy",
		25,
	}
	// reflect.ValueOf(i interface{}).FieldByName(string) 相当于i.name i类型为struct的前提下
	fmt.Printf("%v\n", reflect.ValueOf(c).FieldByName("name"))

	d := new(Foo)
	d.name = "macholy"
	d.age = 25
	fmt.Printf("%v\n", reflect.ValueOf(d))

	//-----------------------
	e := new(Foo)
	// e.name = "deval&macholy"
	// e.age = 25
	e = &Foo{
		"deval&macholy",
		25,
	}
	//不可寻址
	fmt.Printf("%v\n", reflect.ValueOf(*e).FieldByName("name").CanSet())
	// reflect.ValueOf(e).Elem().SetInt(30)
	fmt.Printf("%v\n", reflect.ValueOf(e).Elem())

	str1 := "hello"
	// int1 := 25
	fmt.Printf("%v\n", reflect.TypeOf(str1).String())
	fmt.Printf("%v\n", reflect.TypeOf(str1).Name())

	f := Foo{
		"macholy2",
		30,
	}
	fmt.Printf("%v\n", reflect.TypeOf(f))
	fmt.Printf("%v\n", reflect.TypeOf(f).Name())
	fmt.Printf("%v\n", reflect.TypeOf(f).String())

	g := Foo{}
	h := reflect.TypeOf(g)
	for i := 0; i < h.NumField(); i++ {
		field := h.Field(i)
		fmt.Printf("%v type is %v\n", field.Name, field.Type)
	}
	field2, _ := h.FieldByName("age")
	fmt.Printf("%v \n", field2.Name)
	i := reflect.TypeOf(g)
	// fmt.Printf("%v\n", h.FieldByName("name"))
	field3, _ := i.FieldByName("name")
	fmt.Printf("%v\n", field3.Tag)
	field4 := i.Field(0)
	fmt.Printf("%v\n", field4.Tag)
	fmt.Printf("%v\n", i.NumMethod())
	fmt.Printf("%v\n", i.Method(0))
	fmt.Printf("%v\n", i.Method(0).Name)
	fmt.Printf("%v\n", i.Method(0).Type)

	//----------------------------------------
	//Kind()返回的是基本类型 不是静态类型
	j := Foo{
		"deval2",
		30,
	}
	k := reflect.TypeOf(j)
	fmt.Printf("%v\n", k.Kind())

	//---------------------------------------
	//reflect.Value 类型
	l := Foo{
		"macholy2",
		30,
	}
	m := reflect.ValueOf(l)
	n := reflect.TypeOf(l)
	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", m.Interface() == l)
	fmt.Printf("%v\n", m.FieldByName("name"))
	fmt.Println(reflect.TypeOf(l).FieldByName("name"))
	fmt.Printf("%v\n", m.Field(1))
	fmt.Printf("%v\n", m.NumMethod())
	fmt.Printf("%v\n", m.NumField())
	for i := 0; i < m.NumField(); i++ {
		// fmt.Printf("%v & %v & %v\n", n.Field(i).Name, m.Field(i).Type(), m.Field(i).Interface())
		// fmt.Printf("%s type is :%s ,value is %v\n", n.Field(i).Name, m.Field(i).Type(), m.Field(i).Interface())
		// fmt.Printf("%v\n", m.Field(i))
		tt := n.Field(i)
		vv := m.Field(i)
		fmt.Printf("%v %v\n", tt, vv)
		fmt.Printf("%v %v\n", tt.Name, vv.Type())
	}

	//-------------------------------------
	//设置value的值
	val1 := 10
	val1re := reflect.ValueOf(&val1).Elem()
	fmt.Printf("%v\n", val1re.CanSet())
	val1re.SetInt(456)
	fmt.Printf("%v\n", val1)

	foo1 := &Foo{
		"macholy3",
		30,
	}
	age1 := reflect.ValueOf(foo1).MethodByName("GetAge").Call([]reflect.Value{})
	fmt.Printf("%v\n", age1)

	//------------------------得到tag
	oo := Foo{
		"deval3",
		30,
	}
	ooT := reflect.TypeOf(oo)
	// oov := reflect.ValueOf(oo)
	fmt.Printf("%v\n", ooT.Field(0).Tag)
	// fmt.Printf("%v\n", oov.Field(0).Tag)
}
