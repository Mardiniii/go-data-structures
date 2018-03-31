package main

import "fmt"

// uniqueness check if an element is present in tree or not
func (b *BinaryTree) uniqueValue(val interface{}) bool {
	if b.root == nil {
		return true
	}

	var q []*Node
	q = append(q, b.root)

	for len(q) != 0 {
		n := q[0]
		q = q[1:]

		if n.Value() == val {
			return false
		}

		if n.Left() != nil {
			q = append(q, n.Left())
		}
		if n.Right() != nil {
			q = append(q, n.Right())
		}
	}

	return true
}

// Insert method to add a new val to the tree by level order
// with priority on the left
func (b *BinaryTree) Insert(val interface{}) {
	switch {
	case b.root == nil:
		b.root = &Node{value: val}
		b.size++
		return
	case b.uniqueValue(val) == false:
		fmt.Println(val, "is already present in the Binary Tree")
	default:
		var q []*Node
		q = append(q, b.root)

		for len(q) != 0 {
			node := q[0]
			q = q[1:]

			if node.Left() == nil {
				node.SetLeft(&Node{value: val})
				b.size++
				return
			} else if node.Right() == nil {
				node.SetRight(&Node{value: val})
				b.size++
				return
			}

			if node.Left() != nil {
				q = append(q, node.Left())
			}
			if node.Right() != nil {
				q = append(q, node.Right())
			}
		}
	}
}
