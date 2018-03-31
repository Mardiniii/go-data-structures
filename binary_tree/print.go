package main

import "fmt"

// Print tree values using level order from left to right
func (b *BinaryTree) Print() {
	if b.root == nil {
		fmt.Println("Empty Tree")
		return
	}

	var q []*Node
	q = append(q, b.root)

	for len(q) != 0 {
		n := q[0]
		q = q[1:]

		fmt.Println("Value:", n.Value())

		if n.Left() != nil {
			q = append(q, n.Left())
		}
		if n.Right() != nil {
			q = append(q, n.Right())
		}
	}
}
