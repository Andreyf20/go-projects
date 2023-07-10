package binary_search 

import "testing"

func TestBinarySearch(t *testing.T) {
  got := BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8}, 8)
  wanted := true

  if got != wanted {
    t.Errorf("got %t, wanted %t", got, wanted)
  }

  got = BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8}, 1)
  wanted = true

  if got != wanted {
    t.Errorf("got %t, wanted %t", got, wanted)
  }

  got = BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3)
  wanted = true

  if got != wanted {
    t.Errorf("got %t, wanted %t", got, wanted)
  }

  got = BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8}, 6)
  wanted = true

  if got != wanted {
    t.Errorf("got %t, wanted %t", got, wanted)
  }

  got = BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8}, 22)
  wanted = false 

  if got != wanted {
    t.Errorf("got %t, wanted %t", got, wanted)
  }
}
