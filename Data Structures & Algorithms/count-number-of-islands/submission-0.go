func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	var fill func(x, y int)

	fill = func(x, y int) {
		if x >= m || y >= n || x < 0 || y < 0 {
			// fmt.Printf("Out of bound at %d, %d\n", x, y)
			return
		}

		if grid[x][y] != byte('1') {
			// fmt.Printf("Not land at %d, %d, %c\n", x, y, grid[x][y])
			return
		}

		// fmt.Printf("Filling %d, %d\n", x, y)
		
		grid[x][y] = byte('.')

		fill(x + 1, y)
		fill(x, y + 1)
		fill(x - 1, y)
		fill(x, y - 1)
	}

	count := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] != byte('1') {
				continue
			}
			fill(x, y)
			count += 1
		}
	}

	return count
}
