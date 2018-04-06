package main

// Traverse iterates over all the noodes rooted with the current node
func (n *Node) Traverse(f func(v int)) {
	var i int
	// Find first key greater or equal to given value
	for i = 0; i < len(n.Values()); i++ {
		if !n.IsLeaf() {
			n.ChildAt(i).Traverse(f)
		}
		f(n.ValueAt(i))
	}

	// Since every node have n+1 children we ensure to print
	// the last child of the node
	if !n.IsLeaf() {
		lastChild := n.ChildAt(i)
		lastChild.Traverse(f)
	}
}
