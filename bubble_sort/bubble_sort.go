package bubble_sort

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i += 1 {
		for j := 0; j < len(arr)-1-i; j += 1 {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}
