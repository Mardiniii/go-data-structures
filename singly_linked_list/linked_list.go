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

// Head method to return the head node
func (l *LinkedList) Head() *Node {
	return l.head
}

// Tail method to return the tail node
func (l *LinkedList) Tail() *Node {
	return l.tail
}

// Length method to return the list length
func (l *LinkedList) Length() int {
	return l.length
}

// Setters

// SetHead method to set the head node
func (l *LinkedList) SetHead(n *Node) {
	l.head = n
}

// SetTail method to set the head node
func (l *LinkedList) SetTail(n *Node) {
	l.tail = n
}

// SetLength method to set the head node
func (l *LinkedList) SetLength(length int) {
	l.length = length
}

// Linked list structure methods

// Insert new node at the end of the linked list
func (l *LinkedList) Insert(val interface{}) {
	n := &Node{value: val}

	if l.Head() == nil {
		l.SetHead(n)
	} else {
		l.Tail().SetNext(n)
	}
	l.SetTail(n)
	l.SetLength(l.Length() + 1)
}

// InsertAt method adds a new value at the given position
func (l *LinkedList) InsertAt(pos int, val interface{}) {
	n := &Node{value: val}
	// If the given position is lower than the list length
	// the element will be inserted at the end of the list
	switch {
	case l.Length() < pos:
		l.Insert(val)
	case pos == 1:
		n.SetNext(l.Head())
		l.SetHead(n)
	default:
		node := l.Head()
		// Position - 2 since we want the element replacing the given position
		for i := 1; i < (pos - 1); i++ {
			node = node.Next()
		}
		n.SetNext(node.Next())
		node.SetNext(n)
	}

	l.SetLength(l.Length() + 1)
}

// Get value in the given position
func (l *LinkedList) Get(pos int) interface{} {
	if pos > l.Length() {
		return nil
	}

	node := l.Head()
	// Position - 1 since we want the value in the given position
	for i := 0; i < pos-1; i++ {
		node = node.Next()
	}

	return node.Value()
}

// Delete value at the given position
func (l *LinkedList) Delete(pos int) bool {
	if pos > l.Length() {
		return false
	}

	node := l.Head()
	if pos == 1 {
		l.SetHead(node.Next())
	} else {
		for i := 1; i < pos-1; i++ {
			node = node.Next()
		}
		node.SetNext(node.Next().Next())
	}
	l.SetLength(l.Length() - 1)
	return true
}

// Each method to apply to each element in the linked list a
// function who receives an interface and don't return any value
func (l *LinkedList) Each(f func(val interface{})) {
	for n := l.Head(); n != nil; n = n.Next() {
		f(n.Value())
	}
}

// Print LinkedList
func (l *LinkedList) Print() {
	if l.Head() != nil {
		fmt.Printf("List elements: ")
		for node := l.Head(); node != nil; node = node.Next() {
			fmt.Printf("%+v ", node.Value())
		}
		fmt.Println()
	}
}
