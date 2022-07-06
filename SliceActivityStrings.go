// Shelby Simpson
// Slice Activity: Strings
package main

import (
	"fmt"
	"strings"
)

func main() {
	wordArray := [10]string{}
	belowAvgSlice, aboveAvgSlice := make([]string, 0), make([]string, 0)
	var sum, avg int
	for i := 0; i < 10; i++ {
		fmt.Print("Please enter a word: ")
		fmt.Scan(&wordArray[i])
		wordArray[i] = strings.ToLower(wordArray[i])
		sum += len(wordArray[i])
	}
	avg = sum / 10
	for _, val := range wordArray {
		if len(val) > avg {
			aboveAvgSlice = append(aboveAvgSlice, val)
		} else if len(val) < avg {
			belowAvgSlice = append(belowAvgSlice, val)
		}
	}
	fmt.Println("Word array:", wordArray)
	fmt.Println("Above avg words:", aboveAvgSlice)
	fmt.Println("Below avg words:", belowAvgSlice)
}
