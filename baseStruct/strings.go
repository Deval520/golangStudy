package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "HI,I'M UPPER CASE"

	lower := strings.ToLower(str)

	//check if string contains another string
	if strings.Contains(lower, "upper") {
		fmt.Printf("%s\n", "Yes,exists")
	}

	//strings are array of characters
	//printing out characters 3 to 9
	fmt.Println("Characters 3 to 9:" + str[3:9])

	//printing out first 5 characters
	fmt.Println("First five characters:" + str[:5])

	sentence := "I'm    a sentance made up of words"
	words := strings.Split(sentence, " ")
	fmt.Printf("%v \n", words)
	fmt.Printf("%s \n", words[0])

	//if you were splitting on whitespace, using Fields is better than because
	//it will split on more than just the space,but all whitespace chars
	fields := strings.Fields(sentence)
	fmt.Printf("%v \n", fields)
}
