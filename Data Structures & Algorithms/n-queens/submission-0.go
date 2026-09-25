type pos struct {
	row int
	col int
}

func makeBoard(n int, positions []pos) []string {
	ss := make([]string, n)

	for _, pos := range positions {
		var builder strings.Builder
		for col := range n {
			if col == pos.col {
				builder.WriteString("Q")
			} else {
				builder.WriteString(".")
			}
		}
		ss[pos.row] = builder.String()
	}

	return ss
}

func solveNQueens(n int) [][]string {
	row := map[int]bool{}
	posDiag := map[int]bool{}
	negDiag := map[int]bool{}

	result := [][]pos{}
	current := []pos{}

	var dfs func(c int)
	dfs = func(c int) {
		if c >= n {
			tmp := make([]pos, n)
			copy(tmp, current)
			result = append(result, tmp)
			return
		}

		for r := range n {
			if _, found := row[r]; found {
				continue
			}
			if _, found := posDiag[r - c]; found {
				continue
			}
			if _, found := negDiag[c + r]; found {
				continue
			}

			current = append(current, pos{ col: c, row: r })
			row[r] = true
			posDiag[r - c] = true
			negDiag[c + r] = true

			dfs(c + 1)

			current = current[:len(current) - 1]
			delete(row, r)
			delete(posDiag, r - c)
			delete(negDiag, c + r)
		}
	}

	dfs(0)

	boards := [][]string{}
	for _, r := range result {
		boards = append(boards, makeBoard(n, r))
	}

	// 0 0 - 0 1 - 0 2 - 0 3
	// 1 0 - 1 1 - 1 2 - 1 3
	// 2 0 - 2 1 - 2 2 - 2 3
	// 3 0 - 3 1 - 3 2 - 3 3

	return boards
}
 