package main

// Node struct
type Node struct {
	value interface{}
	left  *Node
	right *Node
}

// Getters

// Left method to return the left node if exists
func (n *Node) Left() *Node {
	return n.left
}

// Right method to return the right node if exists
func (n *Node) Right() *Node {
	return n.right
}

// Value method to return the current node value
func (n *Node) Value() interface{} {
	return n.value
}

// Setters

// SetLeft method to set the left node
func (n *Node) SetLeft(left *Node) {
	n.left = left
}

// SetRight method to set the right node
func (n *Node) SetRight(right *Node) {
	n.right = right
}

// SetValue method to set the value
func (n *Node) SetValue(val interface{}) {
	n.value = val
}

// Node structure methods

// RecursiveInsert method to add a new node with left priority
// func (n *Node) RecursiveInsert(val interface{}) {
// 	if n.left == nil {
// 		n.left = &Node{value: val}
// 	} else if n.right == nil {
// 		n.right = &Node{value: val}
// 	} else {
// 		n.left.RecursiveInsert(val)
// 	}
// }
