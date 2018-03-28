package main

import "fmt"

// Queue struct
type Queue struct {
	End    *Node
	Front  *Node
	Length int
}

// Enqueue method to place a value into the end of the queue
func (q *Queue) Enqueue(val interface{}) {
	newNode := &Node{Value: val}

	if q.Length == 0 {
		q.Front = newNode
		q.End = newNode
	} else {
		currentEnd := q.End
		currentEnd.Previous = newNode

		newNode.Next = currentEnd
		q.End = newNode
	}
	q.Length++
}

// Dequeue method to remove a value into the front of the queue
func (q *Queue) Dequeue() interface{} {
	switch {
	case q.Length == 0:
		return nil
	case q.Length == 1:
		currentFront := q.Front
		q.Front = nil
		q.End = nil
		q.Length--
		return currentFront.Value
	default:
		currentFront := q.Front
		q.Front = currentFront.Previous
		q.Front.Next = nil
		currentFront.Previous = nil
		q.Length--

		return currentFront.Value
	}
}

// Print method to print the content of the current Queue
func (q *Queue) Print() {
	for n := q.Front; n != nil; n = n.Previous {
		fmt.Printf("%+v ", n.Value)
	}
	fmt.Println()
}
