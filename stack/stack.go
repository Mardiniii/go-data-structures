package main

import "fmt"

// Stack struct
type Stack struct {
	Value  interface{}
	Top    *Node
	Length int
}

// Push method to add val at the top of the stack
func (s *Stack) Push(val interface{}) {
	n := &Node{Value: val}

	if s.Top == nil {
		s.Top = n
	} else {
		n.Previous = s.Top
		s.Top = n
	}

	s.Length++
}

// Pop method to remove val at the top of the stack
func (s *Stack) Pop() interface{} {
	if s.Top == nil {
		return nil
	}
	currentTop := s.Top
	s.Top = currentTop.Previous
	currentTop.Previous = nil
	s.Length--

	return currentTop.Value
}

// Print method to print the Stack
func (s *Stack) Print() {
	for n := s.Top; n != nil; n = n.Previous {
		fmt.Println(n.Value)
	}
}
