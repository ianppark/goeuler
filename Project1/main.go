package main

import "fmt"

// UpperLimit - Num to find sum of multiples less than.
const UpperLimit = 1000

func main() {

	sum := 0
	for i := 0; i < UpperLimit; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)

}
