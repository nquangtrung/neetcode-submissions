func findRow(matrix[][]int, target int, top, bottom int) int {
	// fmt.Printf("search for row: %d %d\n", top, bottom)
	if top > bottom {
		return -1;
	}
	mid := (top + bottom) / 2
	// fmt.Printf("mid: %d\n", mid)
	row := matrix[mid]
	if (row[0] <= target && row[len(row) - 1] >= target) {
		// fmt.Printf("mid does contains target\n")
		return mid;
	}

	// fmt.Printf("mid does not contains target\n")	
	if top == bottom {
		// fmt.Printf("top and bottom is the same, end the loop\n")	
		return -1;
	}

	if (row[0] > target) {
		return findRow(matrix, target, top, mid - 1)
	} else {
		return findRow(matrix, target, mid + 1, bottom)
	}
}

func findCell(row[]int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	if row[mid] == target {
		return mid
	}
	if left == right {
		return -1
	}

	if row[mid] < target {
		return findCell(row, target, mid + 1, right)
	} else {
		return findCell(row, target, left, mid - 1)
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}

	row := findRow(matrix, target, 0, len(matrix) - 1)
	// fmt.Printf("found row %d\n", row)
	if row < 0 {
		return false
	}


	cell := findCell(matrix[row], target, 0, len(matrix[row]) - 1)

	return cell >= 0
}
