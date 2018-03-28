// Linked lists implementation in Go, check the next methods
// for the Linked List API:
// insert(value): Insert value at the end of the list
// insertAt(pos, value): Insert value at the given position
// get(pos): Get value at the given position
// delete(pos): Delete value at the given position
// each(func): For loop to traverse the linked list

package main

import "fmt"

// LinkedList struct
type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// Getters

// Length method to return the list length
func (l *LinkedList) Length() int {
	return l.length
}

// Linked list structure methods

// Insert new node at the end of the linked list
func (l *LinkedList) Insert(val interface{}) {
	n := &Node{value: val}

	if l.head == nil {
		l.head = n
	} else {
		l.tail.SetNext(n)
	}
	l.tail = n
	l.length = l.length + 1
}

// InsertAt method adds a new value at the given position
func (l *LinkedList) InsertAt(pos int, val interface{}) {
	n := &Node{value: val}
	// If the given position is lower than the list length
	// the element will be inserted at the end of the list
	switch {
	case l.length < pos:
		l.Insert(val)
	case pos == 1:
		n.SetNext(l.head)
		l.head = n
	default:
		node := l.head
		// Position - 2 since we want the element replacing the given position
		for i := 1; i < (pos - 1); i++ {
			node = node.Next()
		}
		n.SetNext(node.Next())
		node.SetNext(n)
	}

	l.length = l.length + 1
}

// Get value in the given position
func (l *LinkedList) Get(pos int) interface{} {
	if pos > l.length {
		return nil
	}

	node := l.head
	// Position - 1 since we want the value in the given position
	for i := 0; i < pos-1; i++ {
		node = node.Next()
	}

	return node.Value()
}

// Delete value at the given position
func (l *LinkedList) Delete(pos int) bool {
	if pos > l.length {
		return false
	}

	node := l.head
	if pos == 1 {
		l.head = node.Next()
	} else {
		for i := 1; i < pos-1; i++ {
			node = node.Next()
		}
		node.SetNext(node.Next().Next())
	}
	l.length = l.length - 1
	return true
}

// Each method to apply to each element in the linked list a
// function who receives an interface and don't return any value
func (l *LinkedList) Each(f func(val interface{})) {
	for n := l.head; n != nil; n = n.Next() {
		f(n.Value())
	}
}

// Print LinkedList
func (l *LinkedList) Print() {
	if l.head != nil {
		fmt.Printf("List elements: ")
		for node := l.head; node != nil; node = node.Next() {
			fmt.Printf("%+v ", node.Value())
		}
		fmt.Println()
	}
}
