// Shelby Simpson
// Data Collections Activity 3
package main

import (
	"fmt"
	"strings"
)

func multiply(arr []int) int {
	product := 1
	for _, val := range arr {
		product *= val
	}
	return product
}

func add(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	var inputArray = make([]int, 0, 10)
	var temp int
	var anotherInput string
	for i := 0; i < 10; i++ {
		if i >= 5 {
			fmt.Print("Would you like to enter another integer? (y/n): ")
			fmt.Scan(&anotherInput)
			if strings.ToLower(anotherInput) == "n" {
				break
			}
		}
		fmt.Print("Please enter integer ", i, ":")
		fmt.Scan(&temp)
		inputArray = append(inputArray, temp)
	}
	product := multiply(inputArray)
	sum := add(inputArray)
	fmt.Println(inputArray)
	fmt.Println("Product of integers:", product)
	fmt.Println("Sum of integers:", sum)
}
