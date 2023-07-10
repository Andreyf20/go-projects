package bubble_sort

import "testing"

func TestBubbleSort(t *testing.T) {
	got := BubbleSort([]int{1, 2, 3, 4, 5, 6, 7, 8})
	wanted := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for index, value := range got {
		if value != wanted[index] {
			t.Errorf("Wrong Order!!!")
		}
	}

	got = BubbleSort([]int{6, 7, 8, 9, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 0})
	wanted = []int{0, 1, 2, 3, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}

	for index, value := range got {
		if value != wanted[index] {
			t.Errorf("Wrong Order!!!")
		}
	}
}
