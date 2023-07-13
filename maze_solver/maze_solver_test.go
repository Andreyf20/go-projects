package maze_solver

import (
	"testing"
)

func TestMazeSolver(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	mazeResult := []Point{
		Point{X: 10, Y: 0},
		Point{X: 10, Y: 1},
		Point{X: 10, Y: 2},
		Point{X: 10, Y: 3},
		Point{X: 10, Y: 4},
		Point{X: 9, Y: 4},
		Point{X: 8, Y: 4},
		Point{X: 7, Y: 4},
		Point{X: 6, Y: 4},
		Point{X: 5, Y: 4},
		Point{X: 4, Y: 4},
		Point{X: 3, Y: 4},
		Point{X: 2, Y: 4},
		Point{X: 1, Y: 4},
		Point{X: 1, Y: 5},
	}

	got := SolveMaze(maze, "x", Point{10, 0}, Point{1, 5})
	if len(got) == 0 {
		t.Errorf("Error empty result")
	}
	for i := range mazeResult {
		if mazeResult[i] != got[i] {
			t.Errorf("Error different result")
		}
	}
}
