func isPalindrom(s string, l, r int) bool {
	for l <= r {
		if s[l] != s[r] {
			return false
		}

		l += 1
		r -= 1
	}

	return true
}
func partition(s string) [][]string {
	result := [][]string{}
	part := []string{}

	var partition func(i int)
	partition = func(i int) {
		if i >= len(s) {
			tmp := make([]string, len(part))
			copy(tmp, part)
			result = append(result, tmp)
			return
		}

		for j := i; j < len(s); j++ {
			if !isPalindrom(s, i, j) {
				continue
			}
			part = append(part, s[i:j + 1])
			partition(j + 1)
			part = part[:len(part) - 1]
		}
	}

	partition(0)

	return result
}
