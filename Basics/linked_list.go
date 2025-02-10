package main

import "fmt"

/*
	A linked list is a data structure used to store many data.

The basic structure consists of multiple nodes that are linked to each other by pointers.
A node typically consists of stored data and a pointer to the next node.
*/
type Node struct {
	data int
	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Value() int {
	return n.data
}

type LinkedList struct {
	head   *Node
	length int
}

func main() {
	list := LinkedList{nil, 0}

	list.insertAtHead(2)
	list.insertAtTail(3)

	list.print()
}

func (l *LinkedList) insertAtHead(data int) {
	temp1 := &Node{data, nil}

	if l.head == nil {
		l.head = temp1
	} else {
		temp2 := l.head
		l.head = temp1
		temp1.next = temp2
		//l.head, temp1.next = temp1,l.head    -> this works?
	}
	l.length += 1
}

func (l *LinkedList) insertAtTail(data int) {
	temp1 := &Node{data, nil}

	if l.head == nil {
		l.head = temp1
	} else {
		temp2 := l.head
		for temp2.next != nil {
			temp2 = temp2.next
		}
		temp2.next = temp1
	}
	l.length += 1
}

/* func (l *LinkedList) insert(n, data int) {
    if n == 0 {
        l.insertAtHead(data)
    } else if n == l.length-1 {
        l.insertAtTail(data)
    } else {
        temp1 := &Node{data, nil}
        temp2 := l.head

        for i := 0; i < n-1; i++ {
            temp2 = temp2.next
        }
        temp1.next = temp2.next
        temp2.next = temp1
        l.length += 1
    }
} */

func (l *LinkedList) deleteAtHead() {
	temp := l.head
	l.head = temp.next

	l.length -= 1
}

func (l *LinkedList) deleteAtTail() {
	temp1 := l.head
	var temp2 *Node
	for temp1.next != nil {
		temp2 = temp1
		temp1 = temp1.next
	}
	temp2.next = nil

	l.length -= 1
}

func (l *LinkedList) delete(n int) {
	if n == 0 {
		l.deleteAtHead()
	} else if n == l.length-1 {
		l.deleteAtTail()
	} else {
		temp1 := l.head
		for i := 0; i < n-1; i++ {
			temp1 = temp1.next
		}
		temp2 := temp1.next
		temp1.next = temp2.next
		l.length -= 1
	}
}

func (l *LinkedList) Reverse() {
	var curr, prev, next *Node
	curr = l.head
	prev = nil

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func (list *LinkedList) print() {
	var current *Node = list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println()
}
