// Shelby Simpson
// Display Range and Count
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	inputSlice := make([]int, 0, 20)
	fmt.Println("Please enter up to 20 integers (bewteen 1 and 100) separated by spaces:")
	scanner := bufio.NewScanner(os.Stdin)
	var temp string
	for scanner.Scan() {
		fmt.Scan(&temp)
		if temp, err := strconv.Atoi(temp); err == nil {
			inputSlice = append(inputSlice, temp)
		}
	}
	// if ok := scanner.Scan(); ok {
	// 	if temp, err := strconv.Atoi(scanner.Text()); err == nil {
	// 		inputSlice = append(inputSlice, temp)
	// 	}
	// }
	fmt.Println(inputSlice)
	sort.Ints(inputSlice)
	var index int
	for i := 1; i < 100; i += 10 {
		fmt.Printf("%v-%v: ", i, i+9)
		if inputSlice[index] <= i+9 {
			fmt.Print(inputSlice[index])
			index++
		}
		fmt.Println()
	}
}
