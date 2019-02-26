package main

import (
	"fmt"
	"strconv"
)

func reverseString(numString string) string {

	chars := []rune(numString)
	rChars := make([]rune, len(numString))

	for i := len(chars) - 1; i >= 0; i-- {
		rChars[i] = chars[(len(chars)-1)-i]
	}
	return string(rChars)
}

func isPalindrome(numString string) bool {
	if numString == reverseString(numString) {
		return true
	}
	return false
}

func main() {

	maxDigits := 3

	highNumSlice := make([]rune, maxDigits)
	for i := 0; i < maxDigits; i++ {
		highNumSlice[i] = '9'
	}
	highNumString := string(highNumSlice)
	highNum, _ := strconv.Atoi(highNumString)

	var product int
	var maxPal int
	for i := highNum; i > 0; i-- {
		for j := highNum; j > 0; j-- {
			product = i * j
			if isPalindrome(strconv.Itoa(product)) {
				if product > maxPal {
					maxPal = product
				}
			}
		}
	}
	fmt.Println(maxPal)

}
