package main

import "fmt"

func main() {
	max := 4000000
	fibs := [3]int{1, 2, 3}

	sum := 0
	fmt.Println(1)
	for fibs[0] < max {

		fibs[0] = fibs[1]
		fibs[1] = fibs[2]
		fibs[2] = fibs[0] + fibs[1]

		if fibs[0]%2 == 0 {
			sum += fibs[0]
		}

	}
	fmt.Println("Sum of even fibs:", sum)

}
