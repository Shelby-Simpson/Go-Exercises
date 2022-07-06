// Shelby Simpson
// Integer Array Increment with Pointers
package main

import "fmt"

func increment(arrPtr *[]int) {
	for i := range *arrPtr {
		(*arrPtr)[i]++
	}
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr)
	increment(&arr)
	fmt.Println(arr)
}
