package main

func main() {
	leafOne := Node{
		values: []int{1, 2},
		leaf:   true,
	}

	leafTwo := Node{
		values: []int{4, 5, 6},
		leaf:   true,
	}

	leafThree := Node{
		values: []int{40, 50},
		leaf:   true,
	}

	leafFour := Node{
		values: []int{70, 80, 82, 84, 86},
		leaf:   true,
	}

	root := Node{
		values:   []int{3, 30, 60},
		leaf:     false,
		children: []*Node{&leafOne, &leafTwo, &leafThree, &leafFour},
	}

	b := BTree{
		root:  &root,
		order: 3,
	}

	b.Exists(40)
	b.Exists(83)
	b.Exists(0)
	b.Exists(170)
}
