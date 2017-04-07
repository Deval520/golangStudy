package main

import "fmt"

func main() {
	//initializes an empty array
	var empty []int

	//initializes array with values
	arrstr := []string{"abc", "def", "ght", "jkl"}

	empty = append(empty, 123)
	empty = append(empty, 456)
	fmt.Printf("%v \n", empty)

	arrstr = append(arrstr, "jj", "lll")

	fmt.Printf("%v \n", arrstr)

	fmt.Println(len(arrstr))

	fmt.Println(arrstr[5])

	//retrieve a slice of a slice
	arrstr2 := arrstr[0:3]
	fmt.Printf("%v \n", arrstr2)

	fmt.Printf("%v \n", arrstr[:3])

	if eleExits("lll", arrstr) {
		fmt.Println("exists!")
	}
}

func eleExits(s string, a []string) bool {
	for _, val := range a {
		if val == s {
			return true
		}
	}
	return false
}
