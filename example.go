package main

import (
	"fmt"
	"math"
	"time"
)

// func Mix(s, r chan struct{}) {
// 	for {
// 		select {
// 		case <-s:
// 			fmt.Println("Received from Bob")
// 			r <- struct{}{}
// 		case <-r:
// 			fmt.Println("Received from Johan")
// 			s <- struct{}{}
// 		}
// 	}
// }

func main() {
	num := make(chan int)
	go generateNum(num)
	go checkPrime(num)
	time.Sleep(time.Second * 1)
}

func generateNum(num chan<- int) {
	for i := 2; i < 30; i++ {
		num <- i
		// time.Sleep(time.Millisecond)
	}
}

func checkPrime(num <-chan int) {
	for {
		tempNum := <-num
		sqrtNum := math.Sqrt(float64(tempNum))
		prime := true
		for j := 2; j <= int(sqrtNum); j++ {
			if tempNum%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			fmt.Println(tempNum, "is prime.")
		}
		// time.Sleep(time.Millisecond)
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
