package main

import "fmt"

type Node struct {
	Value int
}

type Stack struct {
	nodes []*Node
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	node := s.nodes[0]
	s.nodes = s.nodes[1:]
	return node
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

func main() {
	s := NewStack()
	s.Push(&Node{10})
	s.Push(&Node{12})
	s.Push(&Node{14})
	s.Push(&Node{16})
	fmt.Println(s.Pop(), s.Pop(), s.Pop())
	s.Push(&Node{20})
	fmt.Println(s.Pop())
}
