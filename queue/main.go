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
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Print()
	q.Enqueue(8)
	q.Enqueue(9)
	q.Enqueue(7)
	q.Print()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	fmt.Println(q.Length)
	q.Print()
}
