// Shelby Simpson
// Row and Column Sum
package main

import "fmt"

// Row and Col Sum
func main() {
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	var rowSum, colSum int
	for i := range matrix {
		for j := range matrix[0] {
			rowSum += matrix[i][j]
			colSum += matrix[j][i]
		}
		fmt.Printf("Sum of row %v is %v\n", i, rowSum)
		fmt.Printf("Sum of column %v is %v\n", i, colSum)
		rowSum = 0
		colSum = 0
	}
}
