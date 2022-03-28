package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func InsertAtFirst(head *Node, data int) {
	start := head
	if start == nil {
                 head = &Node{data:data}
		 head.next = nil
	} else {

	}

}

func main() {

	var n, data int
	var Head Node
	fmt.Println("How many number")
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &data)
		InsertAtFirst(&Head, 8)
	}

	fmt.Println("sa", Head)
}
