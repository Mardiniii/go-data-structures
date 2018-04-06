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

// Each receives a function and applies it to every value
// store in the nodes of the tree. This functions makes an
// in order iteration
func (b *BTree) Each(f func(v int)) {
	b.root.Traverse(f)
}
