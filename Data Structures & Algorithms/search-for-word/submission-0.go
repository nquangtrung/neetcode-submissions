func exist(board [][]byte, word string) bool {	
	if len(board) == 0 {
		return false
	}

	picked := map[int]map[int]bool{}

	m, n := len(board), len(board[0])

	var pick func(x, y, current int) bool
	pick = func(x, y, current int) bool {
		if x >= m || y >= n || x < 0 || y < 0 {
			return false
		}

		if picked[x] != nil && picked[x][y] {
			return false
		}

		if word[current] != board[x][y] {
			return false
		}

		if picked[x] == nil {
			picked[x] = map[int]bool{}
		}
		picked[x][y] = true

		if current == len(word) - 1 {
			return true
		}

		if pick(x + 1, y, current + 1) {
			return true
		}

		if pick(x, y + 1, current + 1) {
			return true
		}

		if pick(x - 1, y, current + 1) {
			return true
		}

		if pick(x, y - 1, current + 1) {
			return true
		}
		
		picked[x][y] = false
		return false
	}

	for x := range m {
		for y := range n {
			if pick(x, y, 0) {
				return true
			}
		}
	}

	return false
}
