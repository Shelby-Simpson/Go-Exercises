// Shelby Simpson
// Shifting Array
package main

import "fmt"

// Shift Array
func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var numElements int
	var direction string
	fmt.Println("Please enter the number of elements and direction for the shift")
	fmt.Print("Number of elements {1-8}: ")
	fmt.Scan(&numElements)
DIRECTION:
	fmt.Println("Direction (left or right): ")
	fmt.Scan(&direction)
	if direction == "left" {
		arr1 = append(arr1[numElements:], arr1[0:numElements]...)
	} else if direction == "right" {
		arr1 = append(arr1[len(arr1)-numElements:], arr1[0:len(arr1)-numElements]...)
	} else {
		fmt.Println("You entered an incorrect direction.  Please try again.")
		goto DIRECTION
	}
	fmt.Println(arr1)
}
