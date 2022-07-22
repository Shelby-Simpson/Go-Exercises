package main

import (
	"fmt"
	"time"
)

func goRoutine(ch1 <-chan bool, ch2 chan<- bool) {
	for i := 0; i < 10; i++ {
		select {
		case <-ch1:
			fmt.Println(i)
			ch2 <- true
		}
	}
}

func anotherGoRoutine(ch1 chan<- bool, ch2 <-chan bool) {
	for i := 10; i > 0; i-- {
		select {
		case <-ch2:
			fmt.Println(i)
			ch1 <- true
		}
	}
}

func main() {
	ch1, ch2 := make(chan bool), make(chan bool)
	go goRoutine(ch1, ch2)
	go anotherGoRoutine(ch1, ch2)
	ch1 <- true
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Main func")
}
