package main

import (
	"fmt"
	"math/rand"
)

type tree struct {
	Left  *tree
	Value int
	Right *tree
}

func create(n int) *tree {
	var root *tree
	for _, m := range rand.Perm(10) {
		root = insert(root, (1+m)*n)
	}
	return root
}

func insert(root *tree, val int) *tree {
	if root == nil {
		return &tree{nil, val, nil}
	} else if root.Value < val {
		root.Right = insert(root.Right, val)
		return root
	} else {
		root.Left = insert(root.Left, val)
		return root
	}
}

func walk(root *tree, ch chan int) {
	if root == nil {
		return
	}
	walk(root.Left, ch)
	ch <- root.Value
	walk(root.Right, ch)
}

func same(root1, root2 *tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go walk(root1, ch1)
	go walk(root2, ch2)
	rval := true
	for i := 0; i < 10; i++ {
		val1 := <-ch1
		val2 := <-ch2
		if val1 != val2 {
			rval = false
		}
	}
	return rval
}

func testsame() {
	if same(create(1), create(1)) == true {
		fmt.Println("test case 1 is SUCCEDED!")
	} else {
		fmt.Println("test case 1 is FAILED")
	}
	if same(create(2), create(2)) {
		fmt.Println("test case 2 is SUCCEDED!")
	} else {
		fmt.Println("test case 2 is FAILED")
	}
	if same(create(2), create(1)) {
		fmt.Println("test case 3 is FAILED")
	} else {
		fmt.Println("test case 3 is SUCCEDED!")
	}
}

func main() {
	testsame()
}
