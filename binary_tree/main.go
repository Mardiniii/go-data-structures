package main

import "fmt"

func main() {
	var t BinaryTree
	t.Delete(5)
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
	t.Insert(11)
	t.Insert(12)
	t.Insert(13)
	t.Insert(14)
	t.Insert(15)
	t.Print()
	t.Exists(3)
	t.Exists(9)
	t.Exists(1862534)
	fmt.Println("Tree Size:", t.Size())
	t.Delete(1)
	t.Delete(2)
	t.Delete(3)
	t.Delete(4)
	t.Delete(5)
	t.Delete(1)
	t.Print()
	fmt.Println("Tree Size:", t.Size())

	t.Each(func(val interface{}) {
		fmt.Println("Element:", val)
	})
}
