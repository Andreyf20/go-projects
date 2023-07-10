package queue

import "testing"

func TestQueue(t *testing.T) {
	q := Queue{Head: nil, Tail: nil, Ln: 0}
	Enqueue(&q, 1)
	Enqueue(&q, 2)
	Enqueue(&q, 3)
	Enqueue(&q, 4)

	result := Peek(&q)
	if result != 1 {
		t.Errorf("got %d, wanted %d", result, 1)
	}

	Dequeue(&q)
	result = Peek(&q)
	if result != 2 {
		t.Errorf("got %d, wanted %d", result, 2)
	}
}
