package main

// Node struct
type Node struct {
	value    interface{}
	previous *Node
}

// Getters

// Previous method to return the previous node if exists
func (n *Node) Previous() *Node {
	return n.previous
}

// Value method to return the current node value
func (n *Node) Value() interface{} {
	return n.value
}

// Setters

// SetPrevious method to set the previous node
func (n *Node) SetPrevious(previous *Node) {
	n.previous = previous
}
