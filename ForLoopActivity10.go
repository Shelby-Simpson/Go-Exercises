// Shelby Simpson
// For Loop Activity 10
package main

import "fmt"

// Activity 10
func main() {
	var input int
	fmt.Println("Please enter an integer to construct the pyramid:")
	fmt.Scan(&input)
	for i := 1; i <= input; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Print("\n")
	}
}
