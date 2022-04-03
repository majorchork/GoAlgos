package main

import "fmt"

type LinkedList struct {
	head *Node
}
type Node struct {
	value int
	next  *Node
}

// Append a node to the end of the list => Append
func (l *LinkedList) Append(val int) {
	newNode := &Node{value: val}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Print LinkedList element
func (l *LinkedList) PrintList() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
	fmt.Println()
}

// Prepend Add a node to the beginning of the list
func (l *LinkedList) Prepend(val int) {
	newNode := &Node{value: val}
	if l.head == nil {
		l.head = newNode
	}
	newNode.next = l.head
	l.head = newNode
}

// InsertAt Insert a node after a given node
func (l *LinkedList) InsertAt(val int, pos int) {
	// 10 2 3 : 1 => 10, 15, 2, 3
	newNode := &Node{value: val}
	if pos == 0 {
		newNode.next = l.head
		l.head = newNode
	}
	count := 0
	current := l.head
	for current != nil {
		if count == pos-1 {
			newNode.next = current.next
			current.next = newNode
			break
		}
		count++
		current = current.next
	}
}

// Delete a node from the list => Delete
func (l *LinkedList) Delete(pos int) {
	if l.head == nil {
		return
	}
	if pos == 0 {
		l.head = l.head.next
		return
	}
	count := 0
	current := l.head
	for current != nil {
		if count == pos-1 {
			current.next = current.next.next
		}
		current = current.next
		count++
	}
}
func main() {
	l := LinkedList{}
	l.Append(10)
	l.Append(25)
	l.Append(2)
	l.Prepend(12)
	l.Prepend(24)
	l.InsertAt(35, 3)
	l.InsertAt(10, 6)
	l.InsertAt(11, 8)
	l.Delete(3)
	l.Delete(5)
	// {10, nil}
	// {10, A} {25, nil}
	// {10, A} {25, B} {2, nil}
	l.PrintList()
}
