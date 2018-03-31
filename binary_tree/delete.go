package main

import "fmt"

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
