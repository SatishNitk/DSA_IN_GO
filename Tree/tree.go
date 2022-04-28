package main

import (
	"fmt"
)

type node struct {
	left, right *node
	data        int
}

type Tree struct {
	head *node
}

func (head *node) insert(data int) {

	if head == nil {
		node := &node{data: data}
		head = node
	} else {
		if head.data > data {
			if head.left == nil {

				node := &node{data: data}
				head.left = node
			} else {
				head.left.insert(data)
			}

		} else {
			if head.right == nil {
				node := &node{data: data}
				head.right = node
			} else {
				head.right.insert(data)
			}
		}
	}
}

func (tree *Tree) insert(data int) {

	if tree.head == nil {
		node := &node{data: data}
		tree.head = node
	} else {
		tree.head.insert(data)
	}

}

func inorder(root *node) {

	if root == nil {
		return
	}

	inorder(root.left)
	fmt.Print(root.data, " ")
	inorder(root.right)

}

func main() {

	var tree Tree

	tree.insert(12)

	tree.insert(321)
	tree.insert(22)
	tree.insert(62)
	tree.insert(42)
	tree.insert(2)
	tree.insert(132)
	inorder(tree.head)

}
