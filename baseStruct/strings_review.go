package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "dh dsh jj jk"

	upper := strings.ToUpper(str)

	if strings.Contains(upper, "JJ") {
		fmt.Println("Yes! I'm exists!")
	}

	fmt.Println(str[:1])

	fmt.Println(str[1:5])

	sentence := "dg gh kk ll k"
	words := strings.Split(sentence, " ")

	fmt.Printf("%v\n", words)

	sentenceTwo := "h  jj    k      klk   j"
	wordsTwo := strings.Fields(sentenceTwo)

	fmt.Printf("%v\n", wordsTwo)
}
