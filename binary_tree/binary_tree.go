package main

// BinaryTree struct
type BinaryTree struct {
	root  *Node
	size  int // Number of nodes
	order int // Number of levels
}

// Getters

// Size method to get the size(number of nodes)
func (b *BinaryTree) Size() int {
	return b.size
}

// Order method to get the order(number of levels)
func (b *BinaryTree) Order() int {
	return b.order
}
