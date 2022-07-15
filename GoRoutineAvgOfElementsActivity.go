// Shelby Simpson
// Go Routines to take average of 5 elements each
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	ch := make(chan int)
	go calculateAvg(arr[:5], ch)
	go calculateAvg(arr[5:10], ch)
	go calculateAvg(arr[10:15], ch)
	go calculateAvg(arr[15:], ch)
	avg1, avg2, avg3, avg4 := <-ch, <-ch, <-ch, <-ch
	fmt.Println(avg1)
	fmt.Println(avg2)
	fmt.Println(avg3)
	fmt.Println(avg4)
}

func calculateAvg(arr []int, c chan int) {
	var sum int
	for _, v := range arr {
		sum += v
	}
	c <- sum / 5
}
