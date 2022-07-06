// Shelby Simpson
// For Loop Activity 11
package main

import "fmt"

// Activity 11
func main() {
	var num1, num2 int
	var a, b, r, gcd int
	fmt.Println("Please enter two integers to find the GCD")
	fmt.Println("First integer:")
	fmt.Scan(&num1)
	fmt.Println("Second integer:")
	fmt.Scan(&num2)
	if num1 > num2 {
		a = num1
		b = num2
	} else {
		a = num2
		b = num1
	}
	r = b
	for r != 0 {
		gcd = r
		r = a % b
		a = b
		b = r
	}
	fmt.Printf("The GCD of %v and %v is %v\n", num1, num2, gcd)
}
