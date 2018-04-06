package main

// BTree data structure
type BTree struct {
	root  *Node
	order int
}

// Exists checks for an element in the B-Tree
func (b *BTree) Exists(v int) *Node {
	return b.root.Search(v)
}
