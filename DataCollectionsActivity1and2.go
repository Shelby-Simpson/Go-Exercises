// Shelby Simpson
// Data Collections Activity 1 and 2
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func nSizeArray(n int) []int {
	rand.Seed(time.Now().UnixNano())
	var arr = make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(201) - 100
	}
	return arr
}

func max(arr []int) int {
	var max int = arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func maxIndex(arr []int) int {
	var idx int
	for i, val := range arr {
		if val > arr[idx] {
			idx = i
		}
	}
	return idx
}

func min(arr []int) int {
	var min int = arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
	}
	return min
}

func minIndex(arr []int) int {
	var idx int
	for i, val := range arr {
		if val < arr[idx] {
			idx = i
		}
	}
	return idx
}

func mean(arr []int) int {
	var sum int
	for _, val := range arr {
		sum += val
	}
	return sum / len(arr)
}

func getPositiveSlice(arr []int) []int {
	var positiveNums = make([]int, 0, len(arr))
	for _, val := range arr {
		if val > 0 {
			positiveNums = append(positiveNums, val)
		}
	}
	return positiveNums
}

func getNegativeSlice(arr []int) []int {
	var negativeNums = make([]int, 0, len(arr))
	for _, val := range arr {
		if val < 0 {
			negativeNums = append(negativeNums, val)
		}
	}
	return negativeNums
}

func main() {
	arr := nSizeArray(10)
	fmt.Println(arr)
	fmt.Println(max(arr))
	fmt.Println(maxIndex(arr))
	fmt.Println(min(arr))
	fmt.Println(minIndex(arr))
	fmt.Println(mean(arr))
	fmt.Println(getPositiveSlice(arr))
	fmt.Println(getNegativeSlice(arr))
}
