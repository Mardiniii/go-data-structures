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
	Head   *Node
	Tail   *Node
	Length int
}

// Print LinkedList
func (l *LinkedList) Print() {
	currentNode := l.Head

	if currentNode != nil {
		fmt.Printf("List elements: ")
		for currentNode != nil {
			fmt.Printf("%+v ", currentNode.Value)
			currentNode = currentNode.Next
		}
		fmt.Println()
	}
}

// Insert new node at the end of the linked list
func (l *LinkedList) Insert(val interface{}) {
	n := &Node{Value: val}

	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
	l.Length++
}

// InsertAt method adds a new value at the given position
func (l *LinkedList) InsertAt(pos int, val interface{}) {
	n := &Node{Value: val}
	// If the given position is lower than the list length
	// the element will be inserted at the end of the list
	switch {
	case l.Length < pos:
		l.Insert(val)
	case pos == 1:
		n.Next = l.Head
		l.Head = n
	default:
		currentNode := l.Head
		// Position - 2 since we want the element replacing the given position
		for i := 1; i < (pos - 1); i++ {
			currentNode = currentNode.Next
		}
		n.Next = currentNode.Next
		currentNode.Next = n
	}

	l.Length++
}

// Get value in the given position
func (l *LinkedList) Get(pos int) interface{} {
	if pos > l.Length {
		return nil
	}

	currentNode := l.Head
	// Position - 1 since we want the value in the given position
	for i := 0; i < pos-1; i++ {
		currentNode = currentNode.Next
	}

	return currentNode.Value
}

// Delete value at the given position
func (l *LinkedList) Delete(pos int) bool {
	if pos > l.Length {
		return false
	}

	currentNode := l.Head
	if pos == 1 {
		l.Head = currentNode.Next
	} else {
		for i := 1; i < pos-1; i++ {
			currentNode = currentNode.Next
		}
		currentNode.Next = currentNode.Next.Next
	}
	l.Length--
	return true
}

// Each method to apply to each element in the linked list a
// function who receives an interface and don't return any value
func (l *LinkedList) Each(f func(val interface{})) {
	for n := l.Head; n != nil; n = n.Next {
		f(n.Value)
	}
}
