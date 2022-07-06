// Shelby Simpson
// Data Types Activity 2
package main

import (
	"fmt"
	"math"
)

func main() {
	var floatNum float64
	fmt.Println("Please enter a float64 number: ")
	fmt.Scan(&floatNum)
	fmt.Println(math.Trunc(floatNum))
}
