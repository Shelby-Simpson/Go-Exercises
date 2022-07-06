// Shelby Simpson
// Data Types Activity 7
package main

import "fmt"

func main() {
	var num1, num2, num3, num4, num5 int
	fmt.Println("Please enter 5 integers: ")
	fmt.Println("1: ")
	fmt.Scan(&num1)
	fmt.Println("2: ")
	fmt.Scan(&num2)
	fmt.Println("3: ")
	fmt.Scan(&num3)
	fmt.Println("4: ")
	fmt.Scan(&num4)
	fmt.Println("5: ")
	fmt.Scan(&num5)
	fmt.Printf("You entered the numbers: %v %v %v %v %v\n", num1, num2, num3, num4, num5)
	fmt.Printf("The product of the integers: %v\n", num1*num2*num3*num4*num5)
	fmt.Printf("The average of the integers: %v\n", (num1*num2*num3*num4*num5)/5)
	fmt.Printf("The sum of the integers: %v\n", num1+num2+num3+num4+num5)
}
