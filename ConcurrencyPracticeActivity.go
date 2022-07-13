// Shelby Simpson
// Practice Activity: Concurrency
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		for i := 100; i >= 0; i-- {
			fmt.Println("Anon func 1:", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i <= 100; i++ {
			fmt.Println("Anon func 2:", i)
		}
		wg.Done()
	}()
	wg.Wait()
}
