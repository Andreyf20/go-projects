package maze_solver

type Point struct {
	X int
	Y int
}

var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func Pop(path []Point) []Point {
	return path[:len(path)-1]
}

func Walk(maze []string, wall string, currentPos *Point, end *Point, seen [][]bool, path *[]Point) bool {
	// Invalid spot
	if currentPos.Y < 0 || currentPos.X < 0 || currentPos.Y >= len(maze) || currentPos.X >= len(maze[0]) {
		return false
	}

	// We hit a wall
	// Why this is a byte array even when its type is a string array we'll never know
	if string(maze[currentPos.Y][currentPos.X]) == wall {
		return false
	}

	// We got to the end
	if currentPos.X == end.X && currentPos.Y == end.Y {
		*path = append(*path, *currentPos)
		return true
	}

	// We have already been here
	if seen[currentPos.Y][currentPos.X] {
		return false
	}

	// Pre
	seen[currentPos.Y][currentPos.X] = true
	*path = append(*path, *currentPos)
	// Recurse
	for i := 0; i < len(directions); i += 1 {
		newPoint := Point{currentPos.X + directions[i][0], currentPos.Y + directions[i][1]}
		result := Walk(maze, wall, &newPoint, end, seen, path)
		if result {
			return true
		}
	}

	// Post
	*path = Pop(*path)
	return false
}

func SolveMaze(maze []string, wall string, start Point, end Point) []Point {
	path := []Point{}
	seen := make([][]bool, len(maze))

	for j := 0; j < len(maze); j += 1 {
		for i := 0; i < len(maze[0]); i += 1 {
			seen[j] = append(seen[j], false)
		}
	}

	Walk(maze, wall, &start, &end, seen, &path)
	return path
}
