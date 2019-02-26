package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	factNum := 600851475143

	var maxPrime int
	for i := 2; i < int(math.Sqrt(float64(factNum))); i++ {
		if factNum%i == 0 {
			if isPrime(i) {
				maxPrime = i
			}
		}
	}

	fmt.Println(maxPrime)
}
