package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
	size int
}

func (l *List) append(val int) {
	node := Node{data: val, next: nil}
	if l.head == nil {
		l.head = &node
	} else {
		last := l.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	l.size++
}

func (l *List) insert(pos int, val int) {
	if pos < 0 || pos > l.size {
		panic("Out of index")
	}

}

func (l *List) remove(pos int) {

}

func (l *List) isEmpty() {

}

func (l *List) String() {
	node := l.head
	for {
		if node == nil {
			break
		}
		fmt.Print(node.data)
		fmt.Print("->")
		node = node.next
	}
	fmt.Println()
}

func main() {
	list := List{head: nil}
	list.String()
	list.append(1)
	list.String()
}
