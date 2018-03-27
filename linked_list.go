// Linked lists implementation in Go, check the next methods
// for the Linked List API:
// insert(value): Insert value at the end of the list
// insertAt(pos, value): Insert value at the given position
// get(pos): Get value at the given position
// delete(pos): Delete value at the given position
// each(func): For loop to traverse the linked list

package main

import "fmt"

// Node struct
type Node struct {
	value interface{}
	next  *Node
}

// NextNode function to return the next node if exists
// As an example of how to create a structure method
func (n *Node) NextNode() *Node {
	return n.next
}

// LinkedList struct
type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// Print LinkedList
func (l *LinkedList) print() {
	currentNode := l.head

	if currentNode != nil {
		fmt.Printf("List elements: ")
		for currentNode != nil {
			fmt.Printf("%+v ", currentNode.value)
			currentNode = currentNode.next
		}
		fmt.Println()
	}
}

// Insert new node at the end of the linked list
func (l *LinkedList) insert(n *Node) {
	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
	}
	l.tail = n
	l.length++
}

// Insert new value at the given position
func (l *LinkedList) insertAt(pos int, n *Node) {
	// If the given position is lower than the list length
	// the element will be inserted at the end of the list
	switch {
	case l.length < pos:
		l.insert(n)
	case pos == 1:
		n.next = l.head
		l.head = n
	default:
		currentNode := l.head
		// Position - 2 since we want the element replacing the given position
		for i := 1; i < (pos - 1); i++ {
			currentNode = currentNode.next
		}
		n.next = currentNode.next
		currentNode.next = n
	}

	l.length++
}

// Get value in the given position
func (l *LinkedList) get(pos int) interface{} {
	if pos > l.length {
		return nil
	}

	currentNode := l.head
	// Position - 1 since we want the value in the given position
	for i := 0; i < pos-1; i++ {
		currentNode = currentNode.next
	}

	return currentNode.value
}

// Delete value at the given position
func (l *LinkedList) delete(pos int) bool {
	if pos > l.length {
		return false
	}

	currentNode := l.head
	if pos == 1 {
		l.head = currentNode.next
	} else {
		for i := 1; i < pos-1; i++ {
			currentNode = currentNode.next
		}
		currentNode.next = currentNode.next.next
	}
	l.length--
	return true
}

// Each method to apply function to each element in the linked list
func (l *LinkedList) each() {
	// To be implemented!
}

func main() {
	var l LinkedList
	l.insert(&Node{value: 1})
	l.insert(&Node{value: 2})
	l.insert(&Node{value: 3})
	l.insert(&Node{value: 4})
	l.insert(&Node{value: 5})
	l.insert(&Node{value: 6})
	l.insert(&Node{value: 7})
	l.insert(&Node{value: 8})
	l.insert(&Node{value: 9})
	l.print()
	fmt.Println(l.length)
	l.insertAt(2, &Node{value: "Element"})
	l.insertAt(4, &Node{value: "Element"})
	l.insertAt(6, &Node{value: "Element"})
	l.insertAt(16, &Node{value: "Element"})
	l.print()
	fmt.Println(l.length)
	l.delete(2)
	l.print()
	l.delete(3)
	l.print()
	l.delete(4)
	l.print()
	l.delete(10)
	l.print()
}
