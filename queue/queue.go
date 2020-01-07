package queue

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
	if q.size == 0 {

	}
}

func (q *Queue) dequeue() {
	return
}

func (q *Queue) peek() {

}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}
