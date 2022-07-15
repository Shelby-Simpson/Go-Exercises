// Shelby Simpson
// Go Routine sort slice activity
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	slice := make(chan []int)
	go appendToSlice(slice)
	go sortSlice(slice)
	slice <- []int{1, 2, 3}
	time.Sleep(time.Second * 1)
}

func appendToSlice(slice chan []int) {
	for {
		temp := <-slice
		newTemp := append(temp, rand.Intn(50))
		fmt.Println("In append", newTemp)
		slice <- newTemp
	}
}

func sortSlice(slice chan []int) {
	for i := 0; i < 5; i++ {
		temp := <-slice
		sort.Ints(temp)
		fmt.Print("In sort", temp, " ")
		slice <- temp
	}

}
