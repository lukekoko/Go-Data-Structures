package queue

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type Queue struct {
	front *Node
	back  *Node
	size  int
}

func (q *Queue) new() {
	q.front = nil
	q.back = nil
	q.size = 0
}

func (q *Queue) enqueue(val interface{}) {
	node := Node{data: val, next: nil}
	if q.size == 0 {
		q.front = &node
		q.back = &node
		q.size++
	} else {
		q.back.next = &node
		q.back = &node
		q.size++
	}

}

func (q *Queue) dequeue() *Node {
	if q.size > 0 {
		node := q.front
		q.front = q.front.next
		q.size--
		return node
	}
	return nil
}

func (q *Queue) peek() *Node {
	return q.front
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func (q *Queue) print() {
	node := q.front
	for node != nil {
		fmt.Print(node.data)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}
