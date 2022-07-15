// Shelby Simpson
// Go Routine paragraph activity
package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	sentences := []string{
		"Hello my name is Shelby.",
		"What's your name?",
		"I love to code and you should too.",
		"It's very important that we all know how to code.",
	}
	for _, sentence := range sentences {
		wg.Add(2)
		go printReverseOrder(sentence)
		go countWords(sentence)
		wg.Wait()
	}
	time.Sleep(time.Second * 1)
}

func printReverseOrder(s string) {
	for i := range s {
		fmt.Printf("%c", s[len(s)-i-1])
	}
	fmt.Println()
	time.Sleep(time.Millisecond * 1)
	wg.Done()
}

func countWords(s string) {
	fmt.Println(len(strings.Fields(s)))
	wg.Done()
}
