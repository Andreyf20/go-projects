package two_crystal_balls

import (
	"math"
)

func TwoCrystalBalls(breaks []bool) int {
	n := len(breaks)
	step := int(math.Sqrt(float64(n)))
	i := step
	for ; i < len(breaks); i += step {
		if breaks[i] {
			break
		}
	}

	i -= step
	for i < len(breaks) {
		if breaks[i] {
			return i
		}
    i += 1
	}
	return -1
}
