package quick_sort

func QuickSortAux(arr *[]int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIndex := QuickSortAuxPartition(arr, lo, hi)

	QuickSortAux(arr, lo, pivotIndex-1)
	QuickSortAux(arr, pivotIndex+1, hi)
}

func QuickSortAuxPartition(arr *[]int, lo int, hi int) int {
	pivot := (*arr)[hi]

	idx := lo - 1

	for i := lo; i < hi; i += 1 {
		if (*arr)[i] <= pivot {
			idx += 1
			tmp := (*arr)[i]
			(*arr)[i] = (*arr)[idx]
			(*arr)[idx] = tmp
		}
	}

	idx += 1
	(*arr)[hi] = (*arr)[idx]
	(*arr)[idx] = pivot

	return idx
}

func QuickSort(arr *[]int) {
	if arr == nil || len(*arr) == 1 || len(*arr) == 0 {
		return
	}

	QuickSortAux(arr, 0, len(*arr)-1)
}
