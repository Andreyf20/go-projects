package binary_search

func BinarySearch(haystack []int, needle int) bool {
  left := 0
  right := len(haystack)
  for left < right {
    // This division will have the type of the left and right which are ints so it will be an int
    middle := (left + right) / 2
    if haystack[middle] == needle {
      return true
    }
    if needle < haystack[middle] {
      right = middle
    }
    if needle > haystack[middle] {
      left = middle + 1
    }
  }
  return false
}

