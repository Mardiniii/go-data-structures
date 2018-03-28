package main

// Node struct
type Node struct {
	Value interface{}
	Next  *Node
}

// NextNode function to return the next node if exists
// As an example of how to create a structure method
func (n *Node) NextNode() *Node {
	return n.Next
}
