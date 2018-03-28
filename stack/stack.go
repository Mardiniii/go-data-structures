package main

import "fmt"

// Stack struct
type Stack struct {
	top    *Node
	length int
}

// Getters

// Top method to return the current top node
func (s *Stack) Top() *Node {
	return s.top
}

// Length method to return the stack length
func (s *Stack) Length() int {
	return s.length
}

// Setters

// SetTop method to set the top
func (s *Stack) SetTop(n *Node) {
	s.top = n
}

// SetLength method to set the top
func (s *Stack) SetLength(length int) {
	s.length = length
}

// Stack structure methods

// Push method to add val at the top of the stack
func (s *Stack) Push(val interface{}) {
	n := &Node{value: val}

	if s.Top() == nil {
		s.SetTop(n)
	} else {
		n.SetPrevious(s.Top())
		s.SetTop(n)
	}

	s.SetLength(s.Length() + 1)
}

// Pop method to remove val at the top of the stack
func (s *Stack) Pop() interface{} {
	if s.Top() == nil {
		return nil
	}
	currentTop := s.Top()
	s.SetTop(currentTop.Previous())
	currentTop.SetPrevious(nil)
	s.SetLength(s.Length() - 1)

	return currentTop.Value()
}

// Print method to print the Stack
func (s *Stack) Print() {
	for n := s.Top(); n != nil; n = n.Previous() {
		fmt.Println(n.Value())
	}
}
