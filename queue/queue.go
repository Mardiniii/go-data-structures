package main

import "fmt"

// Queue struct
type Queue struct {
	front  *Node
	end    *Node
	length int
}

// Getters

// Length method to return the current length
func (q *Queue) Length() int {
	return q.length
}

// Queue structure methods

// Enqueue method to place a value into the end of the queue
func (q *Queue) Enqueue(val interface{}) {
	newNode := &Node{value: val}

	if q.length == 0 {
		q.front = newNode
		q.end = newNode
	} else {
		currentEnd := q.end
		currentEnd.SetPrevious(newNode)

		newNode.SetNext(currentEnd)
		q.end = newNode
	}
	q.length = q.length + 1
}

// Dequeue method to remove a value into the front of the queue
func (q *Queue) Dequeue() interface{} {
	switch {
	case q.length == 0:
		return nil
	case q.length == 1:
		currentFront := q.front
		q.front = nil
		q.end = nil
		q.length = q.length - 1
		return currentFront.Value()
	default:
		currentFront := q.front
		q.front = currentFront.Previous()
		q.front.SetNext(nil)
		currentFront.SetPrevious(nil)
		q.length = q.length - 1

		return currentFront.Value()
	}
}

// Print method to print the content of the current Queue
func (q *Queue) Print() {
	for n := q.front; n != nil; n = n.Previous() {
		fmt.Printf("%+v ", n.Value())
	}
	fmt.Println()
}
