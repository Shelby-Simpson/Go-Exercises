// Shelby Simpson
// Go Routine structure channel activity
package main

import "fmt"

func main() {
	send := make(chan int)
	receive := make(chan struct{})

	go printSum(receive, send)
	send <- 1
	receive <- struct{}{}
	<-receive
	fmt.Println("Finished")
}

func printSum(receive chan struct{}, send chan int) {
	for {
		select {
		case num := <-send:
			fmt.Println(num)
		case <-receive:
			receive <- struct{}{}
			return
		}
	}
}
