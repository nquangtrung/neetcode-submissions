func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i, _ := range grid {
		visited[i] = make([]bool, n)
	}

	var fill func(x, y int) int
	fill = func(x, y int) int {
		if x >= m || y >= n || x < 0 || y < 0 {
			return 0
		}

		if grid[x][y] != 1 {
			return 0
		}

		if visited[x][y] {
			return 0
		}

		visited[x][y] = true
		grid[x][y] = -1
		return 1 + fill(x + 1, y) + fill(x, y + 1) + fill(x - 1, y) + fill(x, y - 1)
	}

	maxArea := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] != 1 {
				continue
			}
			maxArea = max(maxArea, fill(x, y))
		}
	}

	return maxArea
}
