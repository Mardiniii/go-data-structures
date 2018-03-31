package main

import "fmt"

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

//Binary Tree structure methods

// uniqueness check if an element is not present yet in the tree
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

// Exists returns an string depending if the value is present or not in the tree
func (b *BinaryTree) Exists(v interface{}) {
	if b.uniqueValue(v) {
		fmt.Printf("The value %v does not exists in the tree!\n", v)
	} else {
		fmt.Printf("The value %v exists in the tree!\n", v)
	}
}

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

// Delete tree value and replaces with deepest rightmost node. Check: https://bit.ly/2uwt0b8
func (b *BinaryTree) Delete(val interface{}) {
	if b.root == nil {
		fmt.Println("Empty Tree")
		return
	}

	var q []*Node
	var dtn *Node // Delete this node
	var n *Node   // Deepest rightmost node
	q = append(q, b.root)

	for len(q) != 0 {
		n = q[0]
		q = q[1:]

		if n.Value() == val {
			dtn = n
		}
		if n.Left() != nil {
			q = append(q, n.Left())
		}
		if n.Right() != nil {
			q = append(q, n.Right())
		}
	}

	if dtn == nil {
		fmt.Println("The value does not exist in the Tree")
		return
	}

	// At this point n is the deepest rightmost node
	b.deleteDeepestRightMostNode(n)
	dtn.SetValue(n.Value())
	b.size--
}

// Delete the deepest rightmost node in the tree
func (b *BinaryTree) deleteDeepestRightMostNode(drn *Node) {
	var q []*Node
	q = append(q, b.root)

	for len(q) != 0 {
		n := q[0]
		q = q[1:]

		if n.Left() != nil {
			if n.Left() == drn {
				n.SetLeft(nil)
				return
			}
			q = append(q, n.Left())
		}
		if n.Right() != nil {
			if n.Right() == drn {
				n.SetRight(nil)
				return
			}
			q = append(q, n.Right())
		}
	}
}
