package linkedList

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) append(val int) {
	node := Node{data: val, next: nil}
	if l.head == nil {
		l.head = &node
	} else {
		last := l.head
		for last.next != nil {
			last = last.next
		}
		last.next = &node
	}
}

func (l *List) insert(pos int, val int) {
	if pos < 0 || pos > l.size() {
		fmt.Println("Out of index")
		return
	}
	newNode := Node{data: val, next: nil}
	node := l.head
	for i := 0; i <= l.size(); i++ {
		if i == pos {
			break
		}
		node = node.next
	}
	newNode.next = node.next
	node.next = &newNode
}

func (l *List) remove(pos int) {
	if pos < 0 || pos > l.size() {
		fmt.Println("Out of index")
		return
	}
	node := l.head
	if pos == 0 {
		l.head = node.next
	} else {
		for i := 1; i <= l.size(); i++ {
			if i == pos {
				break
			}
			node = node.next
		}
		del := node.next
		node.next = del.next
	}
}

func (l *List) search(val int) bool {
	node := l.head
	for node != nil {
		if node.data == val {
			return true
		}
		node = node.next
	}
	return false
}

func (l *List) isEmpty() bool {
	return l.size() == 0
}

func (l *List) size() int {
	i := 0
	node := l.head
	for node != nil {
		i++
		node = node.next
	}
	return i
}

func (l *List) print() {
	node := l.head
	for {
		if node == nil {
			break
		}
		fmt.Print(node.data)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}
