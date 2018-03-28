package main

import "fmt"

func printX(val interface{}) {
	fmt.Println("X")
}

func main() {
	var l LinkedList
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)
	l.Insert(6)
	l.Insert(7)
	l.Insert(8)
	l.Insert(9)
	l.Print()
	fmt.Println(l.Length)
	l.InsertAt(2, "Element")
	l.InsertAt(4, "Element")
	l.InsertAt(6, "Element")
	l.InsertAt(16, "Element")
	l.Print()
	fmt.Println(l.Length)
	l.Delete(2)
	l.Print()
	l.Delete(3)
	l.Print()
	l.Delete(4)
	l.Print()
	l.Delete(10)
	l.Print()
	l.Each(printX)
}
