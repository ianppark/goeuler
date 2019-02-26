package main

import "fmt"

func reverseString(numString string) (string, string) {

	chars := []rune(numString)
	rChars := make([]rune, len(numString))
	for _, v := range chars {
		fmt.Printf("%v\t%c\n", v, v)
	}

	for i := len(chars) - 1; i > 0; i-- {
		rChars[i] = chars[(len(chars)-1)-i]
	}
	return string(chars), string(rChars)
}

func main() {

	x, y := reverseString("1234")
	fmt.Println(x)
	fmt.Println(y)
}
