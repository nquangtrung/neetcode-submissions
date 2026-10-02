type Point struct {
	x, y int
}

func islandsAndTreasure(grid [][]int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return
	}

	m, n := len(grid), len(grid[0])
	queue := []Point{}

	// Step 1: Collect all treasure rooms (sources)
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 0 {
				queue = append(queue, Point{r, c})
			}
		}
	}

	dirs := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// Step 2: Multi-source BFS
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, d := range dirs {
			nx, ny := curr.x+d.x, curr.y+d.y

			// Skip out of bounds or obstacles (-1) or already visited (distance <= new distance)
			if nx < 0 || ny < 0 || nx >= m || ny >= n || grid[nx][ny] != 2147483647 {
				continue
			}

			// Update shortest distance and push to queue
			grid[nx][ny] = grid[curr.x][curr.y] + 1
			queue = append(queue, Point{nx, ny})
		}
	}
}