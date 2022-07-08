package main

import "fmt"

type animal interface {
	breath()
	walk()
}

type human interface {
	animal
	speak()
}

type student struct {
	name string
}

func (s student) breath() {
	fmt.Println("Breath")
}

func (s student) walk() {
	fmt.Println("Walk")
}

func (s student) speak() {
	fmt.Println("My name is", s.name)
}

func main() {
	var s student = student{"Shelby"}
	s.breath()
	s.walk()
	s.speak()
}
