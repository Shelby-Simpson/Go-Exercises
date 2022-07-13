// Shelby Simpson
// Practice Activity: Concurrency
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch1 := make(chan bool, 0)
	ch2 := make(chan int, 0)
	wg.Add(1)
	go func() {
		for {
			if ok := <-ch1; ok {
				fmt.Println("Received:", <-ch2)
			}
		}
	}()
	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println("Sent:", i)
			ch1 <- true
			ch2 <- i
		}
		wg.Done()
	}()
	wg.Wait()
}
