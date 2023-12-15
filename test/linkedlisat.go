package main

import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = l.Tail.Next
}

func main() {
	// Create a new linked list
	myList := List{}

	// Add elements to the linked list
	ListPushBack(&myList, 1)
	ListPushBack(&myList, 2)
	ListPushBack(&myList, "three")
	ListPushBack(&myList,"almost") // You can add elements of any type

	// Display the linked list
	fmt.Println("Linked List:")
	displayList(&myList)
}

// displayList is a helper function to print the elements of the linked list
func displayList(l *List) {
	current := l.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}
