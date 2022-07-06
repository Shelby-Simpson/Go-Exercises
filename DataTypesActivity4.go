// Shelby Simpson
// Data Types Activity 4
package main

import "fmt"

func main() {
	var principal, rate float64
	var days int
	fmt.Println("Please enter the principal amount: ")
	fmt.Scan(&principal)
	fmt.Println("Please enter the rate of interest: ")
	fmt.Scan(&rate)
	fmt.Println("Please enter the number of days for the loan: ")
	fmt.Scan(&days)
	fmt.Printf("The interest is: %v\n", principal*rate*float64(days)/365)
}
