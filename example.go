package main

import (
	"fmt"
)

func createSlice() []int {
	slice := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Printf("Please enter your value for position %v: ", i)
		fmt.Scan(&slice[i])
	}
	return slice
}

func main() {
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = createSlice()
	}
	for i := range matrix {
		for j := 0; j <= i; j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Print("\n")
	}
}
