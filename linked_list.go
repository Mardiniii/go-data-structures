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
func (list *LinkedList) print() {
	currentNode := list.head

	if currentNode != nil {
		fmt.Printf("List elements: ")
		for currentNode != nil {
			fmt.Printf("%+v ", currentNode.value)
			currentNode = currentNode.next
		}
		fmt.Println()
	}
}

// Insert new value at the end of the linked list
func (list *LinkedList) insert(val interface{}) {
	var newNode = Node{
		value: val,
		next:  nil,
	}

	if list.head == nil {
		list.head = &newNode
	} else {
		list.tail.next = &newNode
	}
	list.tail = &newNode
	list.length++
}

// Insert new value at the given position
func (list *LinkedList) insertAt(position int, val interface{}) {
	// If the given position is lower than the list length
	// the element will be inserted at the end of the list
	if list.length < position {
		list.insert(val)
	} else {
		currentNode := list.head
		// Position - 2 since we want the element replacing the given position
		for i := 0; i < (position - 2); i++ {
			currentNode = currentNode.next
		}

		var newNode = Node{
			value: val,
			next:  currentNode.next,
		}

		currentNode.next = &newNode
		list.length++
	}
}

// Get value in the given position
func (list *LinkedList) get(position int) interface{} {
	if position > list.length {
		return nil
	}

	currentNode := list.head
	// Position - 1 since we want the value in the given position
	for i := 0; i < position-1; i++ {
		currentNode = currentNode.next
	}

	return currentNode.value
}

// Delete value at the given position
func (list *LinkedList) delete(position int) bool {
	if position > list.length {
		return false
	}

	currentNode := list.head
	if position == 1 {
		list.head = currentNode.next
	} else {
		for i := 1; i < position-1; i++ {
			currentNode = currentNode.next
		}
		currentNode.next = currentNode.next.next
	}
	list.length--
	return true
}

func main() {
	var list LinkedList
	list.insert(1)
	list.insert(2)
	list.insert(3)
	list.insert(4)
	list.insert(5)
	list.insert(6)
	list.insert(7)
	list.insert(8)
	list.insert(9)
	list.print()
	fmt.Println(list.length)
	list.insertAt(2, 81)
	list.insertAt(4, 81)
	list.insertAt(6, 81)
	list.insertAt(16, 81)
	list.print()
	fmt.Println(list.length)
	fmt.Println(list.get(2))
	fmt.Println(list.get(4))
	fmt.Println(list.get(6))
	list.delete(1)
	list.print()
	list.delete(2)
	list.print()
}
