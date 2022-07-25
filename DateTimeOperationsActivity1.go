// Shelby Simpson
// Date Time Operations Activity 1
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().ISOWeek())
	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Now().YearDay())
	fmt.Println(time.Now().Day())
	fmt.Println(int(time.Now().Weekday()))
}
