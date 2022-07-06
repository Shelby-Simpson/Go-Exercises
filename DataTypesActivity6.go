// Shelby Simpson
// Data Types Activity 6
package main

import "fmt"

func main() {
	var num float64
	var numBool bool
	fmt.Println("Please enter a number: ")
	fmt.Scan(&num)
	fmt.Printf("You selected %v\n", num)
	if num == 0 {
		numBool = false
	} else {
		numBool = true
	}
	fmt.Printf("The boolean of your number is: %v\n", numBool)
	fmt.Printf("The binary equivalent of your number is: %b\n", int(num))
}
