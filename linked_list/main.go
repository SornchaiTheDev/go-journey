package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type NodePtr *Node

type LinkedList struct {
	head NodePtr
}

func (l *LinkedList) AddLast(value int) {

	tmpNode := &Node{value, nil}
	node := l.head

	if l.head == nil {
		l.head = tmpNode
	} else {
		for node.next != nil {
			node = node.next
		}

		node.next = tmpNode
	}
}

func (l *LinkedList) AddFirst(value int) {

	tmpNode := &Node{value, nil}

	if l.head == nil {
		l.head = tmpNode
	} else {

		node := l.head

		for node.next != nil {
			node = node.next
		}

		tmpNode.next = node
		l.head = tmpNode
	}
}

func (l *LinkedList) Print() {

	if l.head == nil {
		fmt.Println("head is nil")
		return
	}

	for node := l.head; node != nil; node = node.next {
		str := fmt.Sprint(node.value)

		if node.next != nil {
			str += " -> "
		}

		fmt.Print(str)
	}
}

func main() {
	link1 := LinkedList{nil}
	link1.AddFirst(2)
	link1.AddFirst(4)
	link1.Print()
}
