// Shelby Simpson
// For Loop Acitivity 6
package main

import "fmt"

func main() {
	var num int
	fmt.Println("Please enter an integer for the multiplication table:")
	fmt.Scan(&num)
	for i := 1; i <= num*num; i++ {
		fmt.Printf("%v * %v = %v\n", num, i, num*i)
	}
}
