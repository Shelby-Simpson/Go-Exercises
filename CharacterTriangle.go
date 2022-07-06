// Shelby Simpson
// Triangle of Characters with Loops
package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%c", j+65)
		}
		fmt.Print("\n")
	}
}
