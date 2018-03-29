package main

import "fmt"

func main() {
	var t BinaryTree
	t.Print()
	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(3)
	t.Insert(4)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.Insert(6)
	t.Insert(7)
	t.Insert(8)
	t.Insert(8)
	t.Insert(9)
	t.Insert(10)
	t.Insert("Hola")
	t.Exists("Hola")
	t.Exists("Hola")
	t.Exists(1862534)
	t.Print()
	fmt.Println(t.Size())
	// t.Insert(3)
	// t.Insert(4)
	// t.Insert(5)
	// t.Insert(6)
	// t.Insert(7)
	// t.Print()
}
