package linked_list

import "testing"

func TestLinkedList(t *testing.T) {
	l := List{Head: nil}
	// Append and Prepend
	Append(&l, 1)
	Append(&l, 2)
	Append(&l, 3)
	Append(&l, 4)
	Append(&l, 5)
	Prepend(&l, 6)
	Prepend(&l, 7)
	Prepend(&l, 8)
	Prepend(&l, 9)
	Prepend(&l, 10)
	PrintList(&l)

	result := Remove(&l, 8)
	PrintList(&l)

	if result != 8 {
		t.Errorf("got %d, wanted %d", result, 8)
	}

	result = RemoveAt(&l, 2)
	PrintList(&l)

	if result != 2 {
		t.Errorf("got %d, wanted %d", result, 2)
	}

	result = InsertAt(&l, 7, 2)
	PrintList(&l)

	if result != 7 {
		t.Errorf("got %d, wanted %d", result, 7)
	}

	result = GetItem(&l, 7)
	PrintList(&l)

	if result != 4 {
		t.Errorf("got %d, wanted %d", result, 4)
	}
}
