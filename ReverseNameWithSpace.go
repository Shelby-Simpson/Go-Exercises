// Shelby Simpson
// Reverse Name with Space
package main

import (
	"fmt"
	"unicode"
)

func main() {
	name := "Robert Pike"
	var reversedName string
	for _, val := range name {
		if unicode.IsSpace(val) {
			fmt.Printf("%v ", reversedName)
			reversedName = ""
		} else {
			reversedName = string(val) + reversedName
		}
	}
	fmt.Printf("%v\n", reversedName)
}
