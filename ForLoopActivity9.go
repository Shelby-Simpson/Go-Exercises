// Shelby Simpson
// For Loop Activity 9
package main

import (
	"fmt"
	"math"
)

// Activity 9
func main() {
	var input, inputCopy, numDigits int
	var firstDigit, lastDigit int
	var digitSum, digitProduct int = 0, 1
	var inputSqrt int
	var isPrime bool = true
	var factorial int = 1

	fmt.Println("Please enter an integer:")
	fmt.Scan(&input)
	inputCopy = input
	lastDigit = inputCopy % 10
	firstDigit = lastDigit
	for inputCopy > 0 {
		numDigits++
		firstDigit = inputCopy % 10
		digitSum += inputCopy % 10
		digitProduct *= inputCopy % 10
		inputCopy /= 10
	}
	inputSqrt = int(math.Sqrt(float64(input)))
	for i := 2; i <= inputSqrt; i++ {
		if input%i == 0 {
			isPrime = false
			break
		}
	}
	for i := 2; i <= input; i++ {
		factorial *= i
	}

	fmt.Printf("Number of digits: %v\n", numDigits)
	fmt.Printf("First digit: %v, last digit: %v\n", firstDigit, lastDigit)
	fmt.Printf("Sum of the digits: %v\n", digitSum)
	fmt.Printf("Product of the digits: %v\n", digitProduct)
	if isPrime {
		fmt.Println("This integer is prime.")
	} else {
		fmt.Println("This integer is not prime.")
	}
	fmt.Printf("The factorial of the integer is %v\n", factorial)
}
