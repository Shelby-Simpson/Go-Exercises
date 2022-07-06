// Shelby Simpson
// Data Types Activity 3
package main

import (
	"fmt"
	"math"
)

func main() {
	var p, r float64
	var n, t int
	fmt.Println("Enter initial deposit: ")
	fmt.Scan(&p)
	fmt.Println("Enter interest rate: ")
	fmt.Scan(&r)
	fmt.Println("Enter number of times per year interest is calculated: ")
	fmt.Scan(&n)
	fmt.Println("Enter years to calculate: ")
	fmt.Scan(&t)
	fmt.Printf("Valus is: %v\n", p*(math.Pow(1+(r/float64(n)), float64(n*t))))
}
