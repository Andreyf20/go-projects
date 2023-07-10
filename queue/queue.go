package queue

type Node struct {
	Val  int
	Next *Node
}

type Queue struct {
	Head *Node
	Tail *Node
	Ln   int
}

func Enqueue(q *Queue, value int) int {
	newNode := Node{Val: value, Next: nil}
	if q.Tail == nil {
		q.Head = &newNode
		q.Tail = &newNode
	} else {
		q.Tail.Next = &newNode
		q.Tail = &newNode
	}
	q.Ln = q.Ln + 1
	return value
}

func Dequeue(q *Queue) int {
	if q.Head == nil {
		return -1
	}

	currentHead := q.Head
	q.Head = currentHead.Next
	q.Ln = q.Ln - 1
	return currentHead.Val
}

func Peek(q *Queue) int {
	if q.Head != nil {
		return q.Head.Val
	}
	return -1
}
