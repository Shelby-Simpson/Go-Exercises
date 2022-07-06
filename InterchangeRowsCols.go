// Shelby Simpson
// Interchange first and last rows and first and last columns
package main

import "fmt"

func interchangeRows(matrix *[][]int) {
	var temp, matrixLength int = 0, len(*matrix)

	for i := 0; i < matrixLength; i++ {
		temp = (*matrix)[0][i]
		(*matrix)[0][i] = (*matrix)[matrixLength-1][i]
		(*matrix)[matrixLength-1][i] = temp
	}
}

func interchangeCols(matrix *[][]int) {
	var temp, matrixLength int = 0, len(*matrix)

	for i := 0; i < matrixLength; i++ {
		temp = (*matrix)[i][0]
		(*matrix)[i][0] = (*matrix)[i][matrixLength-1]
		(*matrix)[i][matrixLength-1] = temp
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	interchangeCols(&matrix)
	fmt.Println(matrix)
}
