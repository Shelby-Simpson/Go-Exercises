// Shelby Simpson 6/30/22
// Fizz Buzz Assessment
package main

import "fmt"

func main() {
	var input, count int
	i := 1
	fmt.Print("How many fizzing and buzzing units do you need in your life? ")
	fmt.Scan(&input)
	fmt.Println(0)
	for count < input {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
			count++
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Println("Buzz")
			count++
		} else if i%3 == 0 && i%5 != 0 {
			fmt.Println("Fizz")
			count++
		} else {
			fmt.Println(i)
		}
		i++
	}
	fmt.Println("TRADITION!!")
}
