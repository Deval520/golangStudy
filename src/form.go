package main

import "fmt"

type foo struct {
	int
}

type bar foo

func main() {
	str := "hello im a string"

	mystr := []byte(str)

	mystr[0] = byte('H')

	s := bar{1}

	var f foo = foo(s)

	str = string(mystr)

	fmt.Printf("%v %v %v\n", mystr, len(str), len(mystr))

	fmt.Printf("%v\n", str)

	fmt.Printf("%v\n", byte('h'))

	fmt.Printf("%v\n", s)

	fmt.Printf("%v\n", f)

}
