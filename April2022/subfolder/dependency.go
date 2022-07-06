package subfolder

import "fmt"

func DemoFunction() {
	fmt.Println("I am from outside of main.")
}

func demoFunction() {
	fmt.Println("Not callable from inside of main.")
}
