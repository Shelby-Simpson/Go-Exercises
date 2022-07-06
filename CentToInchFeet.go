// Shelby Simpson
// Converting Centimeters to Inches and Feet
package main

import (
	"fmt"
)

func main() {
	var centimeters, inches float32
	var feet int
	fmt.Println("Please enter an amount of centimeters: ")
	fmt.Scan(&centimeters)
	inches = centimeters / 2.54
	feet = int(inches / 12)
	inches -= float32(12 * feet)
	fmt.Printf("Conerted to: %v feet and %v inches\n", feet, inches)
}
