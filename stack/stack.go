package stack

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) new() {
	s.head = nil
	s.size = 0
}

func (s *Stack) push(val interface{}) {
	node := Node{data: val, next: s.head}
	s.head = &node
	s.size++
}

func (s *Stack) pop() *Node {
	if s.head == nil {
		return nil
	}
	node := s.head
	s.head = s.head.next
	s.size--
	return node
}

func (s *Stack) peek() *Node {
	return s.head
}

func (s *Stack) isEmpty() bool {
	return s.size == 0
}

func (s *Stack) print() {
	node := s.head
	for node != nil {
		fmt.Print(node.data)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}
