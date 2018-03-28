package main

import "fmt"

// Stack struct
type Stack struct {
	top    *Node
	length int
}

// Getters

// Length method to return the stack length
func (s *Stack) Length() int {
	return s.length
}

// Stack structure methods

// Push method to add val at the top of the stack
func (s *Stack) Push(val interface{}) {
	n := &Node{value: val}

	if s.top == nil {
		s.top = n
	} else {
		n.SetPrevious(s.top)
		s.top = n
	}

	s.length = s.length + 1
}

// Pop method to remove val at the top of the stack
func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	currentTop := s.top
	s.top = currentTop.Previous()
	currentTop.SetPrevious(nil)
	s.length = s.length - 1

	return currentTop.Value()
}

// Print method to print the Stack
func (s *Stack) Print() {
	for n := s.top; n != nil; n = n.Previous() {
		fmt.Println(n.Value())
	}
}
