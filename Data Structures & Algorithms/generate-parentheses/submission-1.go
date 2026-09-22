func generateParenthesis(n int) []string {
	result := []string{}

	var pick func(current string, o, c int)
	pick = func(current string, o, c int) {
		if o == n && o != c {
			pick(current + ")", o, c + 1)
			return
		}

		if o == c && o == n {
			result = append(result, current)
			return
		}

		if o == c {
			// next one can only be open
			pick(current + "(", o + 1, c)
			return
		}

		pick(current + "(", o + 1, c)
		pick(current + ")", o, c + 1)
	}

	pick("", 0, 0)

	return result
}
