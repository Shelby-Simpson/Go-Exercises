// Shelby Simpson
// Elecricity Bill Switch Statement Exercise
package main

import "fmt"

func main() {
	var units, bill int
	fmt.Println("Please enter the number of units: ")
	fmt.Scan(&units)
	switch {
	case units >= 0 && units <= 100:
		bill = units * 5
	case units >= 101 && units <= 200:
		bill = ((units - 100) * 7) + 500
	case units >= 201 && units <= 350:
		bill = ((units - 200) * 10) + 1200
	default:
		bill = 0
	}
	fmt.Printf("Your bill is: $%v\n", bill)
}
