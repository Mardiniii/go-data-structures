package main

import "fmt"

func main() {
	var s Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Print()
	fmt.Println("Stack Length:", s.Length())

	s.Pop()
	s.Pop()
	s.Pop()

	s.Print()
	fmt.Println("Stack Length:", s.Length())
}
