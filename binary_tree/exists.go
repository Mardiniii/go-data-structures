package main

import "fmt"

// Exists returns an string depending if the value is present or not in the tree
func (b *BinaryTree) Exists(v interface{}) {
	if b.uniqueValue(v) {
		fmt.Printf("The value %v does not exists in the tree!\n", v)
	} else {
		fmt.Printf("The value %v exists in the tree!\n", v)
	}
}
