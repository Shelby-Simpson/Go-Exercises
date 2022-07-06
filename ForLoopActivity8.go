// Shelby Simpson
// For Loop Activity 8
package main

import (
	"fmt"
	"math"
)

func main() {
	var exp float64
	var n, factorial int
	var x int = 3
	fmt.Println("Please enter an integer for the series calculation of e^3:")
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		factorial = 1
		for j := 1; j <= i; j++ {
			factorial *= j
		}
		exp += math.Pow(float64(x), float64(i)) / float64(factorial)
	}
	fmt.Println(exp)
}
