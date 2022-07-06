// Shelby Simpson
// Twin Primes between 1 and 100
package main

import (
	"fmt"
	"math"
)

func main() {
	var numSqrt int
	var prevIsPrime, isPrime bool = false, true
	for i := 3; i < 100; i += 2 {
		// check if number is prime
		numSqrt = int(math.Sqrt(float64(i)))
		isPrime = true
		for j := 2; j <= numSqrt; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime && prevIsPrime {
			fmt.Printf("%v %v\n", i-2, i)
		}
		if isPrime {
			prevIsPrime = true
		} else {
			prevIsPrime = false
		}
	}
}
