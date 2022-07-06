// Shelby Simpson
// Flow Control Practice 2 Else If
package main

import "fmt"

func main() {
	grade := 85
	if grade >= 90 {
		fmt.Println("A")
	} else if grade >= 80 {
		fmt.Println("B")
	} else if grade >= 70 {
		fmt.Println("C")
	} else if grade >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
