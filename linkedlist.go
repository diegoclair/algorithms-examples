package main

import "fmt"

//Linked list consists basically in value and a pointer to the next node
//We also have a Doubly linked list and it basically add a pointer to a previous node too

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

//prepend - add a node to the begining of the linked list
func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) deleteWithValue(v int) {
	if l.length == 0 {
		return
	}
	if l.head.data == v {
		l.head = l.head.next
		l.length--
		return
	}
	previosToDelete := l.head
	for previosToDelete.next.data != v {
		if previosToDelete.next.next == nil {
			return
		}
		previosToDelete = previosToDelete.next
	}
	previosToDelete.next = previosToDelete.next.next
	l.length--
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}
	node4 := &node{data: 4}
	node5 := &node{data: 11}
	node6 := &node{data: 2}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.printListData()
	mylist.deleteWithValue(2)
	mylist.printListData()
	mylist.deleteWithValue(100)
	mylist.printListData()
}
