// Shelby Simpson
// Array Practice 2 Reverse Numbers
package main

import "fmt"

// Array Practice 2
func main() {
	arr1 := [3]int{5, 6, 7}
	arr2 := [3]int{}
	for idx, val := range arr1 {
		arr2[len(arr1)-1-idx] = val
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
}
