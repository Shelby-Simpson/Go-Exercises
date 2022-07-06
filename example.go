package main

import "fmt"

type class struct {
	ClassName string
	students  []student
}

type student struct {
	Name   string
	RollNo int
	City   string
}

func main() {
	abs := student{Name: "Ross", RollNo: 30, City: "New York"}
	cbs := student{Name: "Marry", RollNo: 31, City: "London"}
	students := []student{abs, cbs, {Name: "Jack", RollNo: 32, City: "London"}}
	students = append(students, student{Name: "Kate", RollNo: 33, City: "New York"})
	// fmt.Println(students)
	// Slice of slice of structure
	class := class{"firstA", students}
	fmt.Println(class)
}
