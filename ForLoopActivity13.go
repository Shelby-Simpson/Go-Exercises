// Shelby Simpson
// For Loop Activity 13
package main

import (
	"fmt"
	"math"
)

func main() {
	// For every prime number < input, check if input - pm is prime
	var input, numSqrt int
	var isPrime bool = true
	fmt.Println("Please enter an integer:")
	fmt.Scan(&input)
	for i := 2; i < input; i++ {
		numSqrt = int(math.Sqrt(float64(i)))
		for j := 2; j <= numSqrt; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
	}
}
