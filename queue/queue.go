package main

import "fmt"

// Queue struct
type Queue struct {
	front  *Node
	end    *Node
	length int
}

// Getters

// Front method to return the current front node
func (q *Queue) Front() *Node {
	return q.front
}

// End method to return the current end node
func (q *Queue) End() *Node {
	return q.end
}

// Length method to return the current length
func (q *Queue) Length() int {
	return q.length
}

// Setters

// SetFront method to set the front node
func (q *Queue) SetFront(n *Node) {
	q.front = n
}

// SetEnd method to set the end node
func (q *Queue) SetEnd(n *Node) {
	q.end = n
}

// SetLength method to set the length
func (q *Queue) SetLength(l int) {
	q.length = l
}

// Queue structure methods

// Enqueue method to place a value into the end of the queue
func (q *Queue) Enqueue(val interface{}) {
	newNode := &Node{value: val}

	if q.Length() == 0 {
		q.SetFront(newNode)
		q.SetEnd(newNode)
	} else {
		currentEnd := q.End()
		currentEnd.SetPrevious(newNode)

		newNode.SetNext(currentEnd)
		q.SetEnd(newNode)
	}
	q.SetLength(q.Length() + 1)
}

// Dequeue method to remove a value into the front of the queue
func (q *Queue) Dequeue() interface{} {
	switch {
	case q.Length() == 0:
		return nil
	case q.Length() == 1:
		currentFront := q.Front()
		q.SetFront(nil)
		q.SetEnd(nil)
		q.SetLength(q.Length() - 1)
		return currentFront.Value()
	default:
		currentFront := q.Front()
		q.SetFront(currentFront.Previous())
		q.Front().SetNext(nil)
		currentFront.SetPrevious(nil)
		q.SetLength(q.Length() - 1)

		return currentFront.Value()
	}
}

// Print method to print the content of the current Queue
func (q *Queue) Print() {
	for n := q.Front(); n != nil; n = n.Previous() {
		fmt.Printf("%+v ", n.Value())
	}
	fmt.Println()
}
