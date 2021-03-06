package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
}

func (linkedList *LinkedList) insert(data int) {
	head := linkedList.head

	newNode := &node{data: data}
	if head == nil {
		linkedList.head = newNode
	} else {
		for head.next != nil {

			head = head.next
		}
		head.next = newNode
	}
}

func (linkedList LinkedList) display() {

	head := linkedList.head

	for head != nil {

		fmt.Print(head.data, " ")
		head = head.next
	}
}

func (linkedList LinkedList) findMiddle() *node {

	head := linkedList.head

	slow := head
	fast := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow

}

func (linkedList LinkedList) findCycle() *node {

	head := linkedList.head

	slow := head
	fast := head

	for slow != nil && fast != nil && fast.next != nil {

		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return slow

		} else {
			return nil
		}

	}
	return nil

}

func reverseLinkedList(head *node) *node {

	var prev, curr, next *node

	prev = nil
	curr = head
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next

	}
	return prev

}

func printReverseList(head *node) {
	start := head
	for start != nil {

		fmt.Println(" ", start.data)
		start = start.next
	}
}

func mergeSortedList(first *node, second *node) *node {
	if first == nil {
		return first
	}
	if second == nil {
		return second
	}

}

func main() {

	var linkedList LinkedList

	linkedList.insert(10)
	linkedList.insert(1320)
	linkedList.insert(13220)
	linkedList.insert(101)

	linkedList.display()
	fmt.Println("Middle Element:- ", linkedList.findMiddle().data)
	reverseLinkedList := reverseLinkedList(linkedList.head)
	printReverseList(reverseLinkedList)
}
