package main

import "fmt"

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
