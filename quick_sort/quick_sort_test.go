package quick_sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}

	QuickSort(&arr)

	fmt.Println(arr)

	for i := 0; i < len(arr)-2; i += 1 {
		if arr[i] > arr[i+1] {
			t.Errorf("The list is not sorted properly, index error %d", i)
		}
	}
}
