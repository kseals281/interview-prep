package main

import (
	"fmt"
	"math/rand"
	"time"
)

type leaf struct {
	num   int
	left  *leaf
	right *leaf
}

func setupTree() *leaf {
	rand.Seed(time.Now().Unix())
	head := insert(nil, rand.Intn(100))
	for i := 0; i < 19; i++ {
		insert(head, rand.Intn(100))
	}
	return head
}

func build123() {
	head := new(leaf)
	head.num = 2
	head.left = new(leaf)
	head.left.num = 1
	head.right = new(leaf)
	head.right.num = 3
	inOrder(head)
}

func insert(parent *leaf, value int) *leaf {
	if parent == nil {
		curr := new(leaf)
		curr.num = value
		return curr
	}
	if parent.num >= value {
		parent.left = insert(parent.left, value)
	} else {
		parent.right = insert(parent.right, value)
	}
	return parent
}

func inOrder(parent *leaf) {
	if parent == nil {
		return
	}
	inOrder(parent.left)
	fmt.Print(parent.num)
	fmt.Print(" ")
	inOrder(parent.right)
}

func preOrder(parent *leaf) {
	if parent == nil {
		return
	}
	fmt.Print(parent.num)
	fmt.Print(" ")
	preOrder(parent.left)
	preOrder(parent.right)
}

func postOrder(parent *leaf) {
	if parent == nil {
		return
	}
	postOrder(parent.left)
	postOrder(parent.right)
	fmt.Print(parent.num)
	fmt.Print(" ")
}

func exists(curr *leaf, value int) bool {
	if curr == nil {
		return false
	}
	if value == curr.num {
		return true
	}
	return exists(curr.left, value) || exists(curr.right, value)

}

func size(curr *leaf) int {
	left, right := 0, 0
	if curr.left != nil {
		left = size(curr.left)
	}
	if curr.right != nil {
		right = size(curr.right)
	}
	return left + right + 1
}

func maxDepth(curr *leaf) int {
	if curr == nil {
		return 0
	}
	leftDepth := maxDepth(curr.left)
	rightDepth := maxDepth(curr.right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func minValue(curr *leaf) int {
	for curr.left != nil {
		curr = curr.left
	}
	return curr.num
}
