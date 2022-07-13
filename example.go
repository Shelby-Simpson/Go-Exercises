package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	for n := range c {
		fmt.Println(n)
	}
}

// var wg sync.WaitGroup
// var mutex sync.Mutex
// var Counter int

// func main() {
// 	wg.Add(2)
// 	go increment("Foo")
// 	go increment("Bar")
// 	wg.Wait()
// 	// fmt.Println("Main")
// 	// time.Sleep(1 * time.Second)
// }

// func increment(s string) {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 20; i++ {
// 		mutex.Lock()
// 		Counter++
// 		fmt.Println(s, i, "Counter", Counter)
// 		mutex.Unlock()
// 	}
// 	wg.Done()
// }

// func foo() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("Foo:", i)
// 		// time.Sleep(1 * time.Second)
// 	}
// 	wg.Done()
// }

// func bar() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("Bar:", i)
// 		// time.Sleep(1 * time.Second)
// 	}
// 	wg.Done()
// }
