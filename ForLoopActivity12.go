// Shelby Simpson
// For Loop Activity 12
package main

import (
	"fmt"
	"math"
)

// Activity 12
func main() {
	var num1, num2, a, b, lcd int
	fmt.Println("Please enter two integers to find the LCD")
	fmt.Println("First integer:")
	fmt.Scan(&num1)
	fmt.Println("Second integer:")
	fmt.Scan(&num2)
	num1 = int(math.Abs(float64(num1)))
	num2 = int(math.Abs(float64(num2)))
	if num1 != 0 && num2 != 0 {
		if num1 > num2 {
			a = num1
			b = num2
		} else {
			a = num2
			b = num1
		}
		lcd = a
		for true {
			if lcd%b == 0 {
				break
			} else {
				lcd += a
			}
		}
	}
	fmt.Printf("The LCD of %v and %v is %v\n", num1, num2, lcd)
}
