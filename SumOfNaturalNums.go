// Shelby Simpson
// Calculating the Sum of the First N Natural Numbers
package main

import "fmt"

func main() {
	n := 7
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
}
