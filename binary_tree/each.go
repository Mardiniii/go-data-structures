package main

import "fmt"

// Each method to apply to each element in the tree afunction
// who receives an interface and don't return any value
func (b *BinaryTree) Each(f func(val interface{})) {
	if b.root == nil {
		fmt.Println("Empty Tree")
	}

	var q []*Node
	q = append(q, b.root)

	for len(q) != 0 {
		n := q[0]
		q = q[1:]

		f(n.Value())

		if n.Left() != nil {
			q = append(q, n.Left())
		}
		if n.Right() != nil {
			q = append(q, n.Right())
		}
	}
}
