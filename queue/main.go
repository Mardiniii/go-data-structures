package main

import "fmt"

func main() {
	var q Queue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	q.Enqueue(7)
	q.Print()
	fmt.Println("Queue length:", q.Length())
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Print()
	fmt.Println("Queue length:", q.Length())
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")
	q.Print()
	fmt.Println("Queue length:", q.Length())
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Print()
	fmt.Println("Queue length:", q.Length())
}
