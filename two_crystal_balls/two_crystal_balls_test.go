package two_crystal_balls

import "testing"

func TestLinearSearch(t *testing.T) {
  got := TwoCrystalBalls([]bool{false, false, false, false, false, false, false, false, false, false, true})
  wanted := 10

  if got != wanted {
    t.Errorf("got %d, wanted %d", got, wanted)
  }

  got = TwoCrystalBalls([]bool{false, false, false, false, false, true, true, true, true, true, true})
  wanted = 5

  if got != wanted {
    t.Errorf("got %d, wanted %d", got, wanted)
  }

  got = TwoCrystalBalls([]bool{false, false, false, false, false, false, false, false, false, false, false})
  wanted = -1

  if got != wanted {
    t.Errorf("got %d, wanted %d", got, wanted)
  }
}
