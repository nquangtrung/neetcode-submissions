func getLetters(digit byte) []byte {
	switch digit {
	case '2':
		return []byte{'a', 'b', 'c'}
	case '3':
		return []byte{'d', 'e', 'f'}
	case '4':
		return []byte{'g', 'h', 'i'}
	case '5':
		return []byte{'j', 'k', 'l'}
	case '6':
		return []byte{'m', 'n', 'o'}
	case '7':
		return []byte{'p', 'q', 'r', 's'}
	case '8':
		return []byte{'t', 'u', 'v'}
	case '9':
		return []byte{'w', 'x', 'y', 'z'}
	default:
		return []byte{}
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	result := []string{}

	var dfs func (i int, s string)
	dfs = func(i int, s string) {
		if i >= len(digits) {
			result = append(result, s)
			return
		}

		letters := getLetters(digits[i])
		for _, letter := range letters {
			dfs(i + 1, s + string([]byte{letter}))
		}
	}

	dfs(0, "")

	return result	
}
