type NumberPosition struct {
	row []bool
	col []bool
	sqr []bool
}

func sqrFromPos(x int, y int) int {
	return x / 3 * 3 + y / 3
}

func isValidSudoku(board [][]byte) bool {
	validity := make(map[byte]NumberPosition)

	for x, row := range board {
		for y, cell := range row {
			if cell == byte('.') {
				continue
			}
			sqr := sqrFromPos(x, y)
			// fmt.Printf("for x, y %d, %d => sqr %d\n", x, y, sqr)
			if v, ok := validity[cell]; ok {
				if v.row[x] || v.col[y] || v.sqr[sqr] {
					// fmt.Printf("found %d at %d, %d (%t, %t, %t)", cell - byte('0'), x, y, v.row[x], v.col[y], v.sqr[sqr])
					return false
				}

				validity[cell].row[x] = true
				validity[cell].col[y] = true
				validity[cell].sqr[sqr] = true
				continue
			}

			validity[cell] = NumberPosition{
				row: make([]bool, 9),
				col: make([]bool, 9),
				sqr: make([]bool, 9),
			}
			validity[cell].row[x] = true
			validity[cell].col[y] = true
			validity[cell].sqr[sqr] = true
		}
	}

	return true
}
