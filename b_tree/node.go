package main

import (
	"fmt"
)

// Node struct
type Node struct {
	values   []int
	children []*Node
	leaf     bool
}

// Getters

// IsLeaf method returns true if the node is a leaf node
func (n *Node) IsLeaf() bool {
	return n.leaf
}

// Values method returns values stored in the node
func (n *Node) Values() []int {
	return n.values
}

// ValueAt method returns the value in the given index
func (n *Node) ValueAt(i int) int {
	return n.values[i]
}

// Children method returns the array of children
func (n *Node) Children() []*Node {
	return n.children
}

// ChildAt method returns the child in the given index
func (n *Node) ChildAt(i int) *Node {
	return n.children[i]
}

// Node exported methods

// Search an element in the current node
func (n *Node) Search(v int) *Node {
	index := 0

	// Find first key greater or equal to given value
	for i := 0; i < len(n.Values()); i++ {
		if v > n.ValueAt(i) {
			index++
		}
	}

	if index < len(n.Values()) && v == n.ValueAt(index) {
		fmt.Println("Value:", v, "found in the tree")
		return n
	}

	if n.IsLeaf() == true {
		fmt.Println("Value:", v, "not found in the tree")
		return nil
	}

	return n.ChildAt(index).Search(v)
}
