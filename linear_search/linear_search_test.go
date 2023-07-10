package linear_search

import "testing"

func TestLinearSearch(t *testing.T) {
  got := LinearSearch([]int{1, 2, 3, 4, 5, 6, 7, 8}, 8)
  wanted := true

  if got != wanted {
    t.Errorf("got %t, wanted %t", got, wanted)
  }
}
