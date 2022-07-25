package main

import (
	"fmt"
	"os"
)

func main() {
	var numChars int
	fmt.Print("Please enter number of characters to read from offset: ")
	fmt.Scan(&numChars)
	f, err := os.Open("/Users/shelbysimpson/Documents/golang/fileio/flatland1.txt")
	fmt.Println(err)

	// skip the first 100 bytes
	s, err := f.Seek(-100, 2)
	fmt.Println(err) // nil if no isssues

	// display the offset
	fmt.Println(s)

	// read 5 bytes starting from the offset
	data := make([]byte, numChars)
	n, err := f.Read(data)

	fmt.Println(err) // nil if no issues
	fmt.Println("Bytes read", n)
	fmt.Println("Reading starting from byte", s, ":", string(data[:n]))

	// close the file
	f.Close()
}
